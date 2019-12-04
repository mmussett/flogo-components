package ems

import (
	"encoding/json"
	"errors"
	"fmt"

	ems "github.com/mmussett/ems"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	opentracing "github.com/opentracing/opentracing-go"
	ctx "golang.org/x/net/context"
)

const (
	ivContent         = "content"
	ivDestination     = "destination"
	ivDestinationType = "destinationType"
	ivServerURL       = "serverUrl"
	ivUser            = "user"
	ivPassword        = "password"
	ivDeliveryDelay   = "deliveryDelay"
	ivDeliveryMode    = "deliveryMode"
	ivExpiration      = "expiration"
	ivTracing         = "tracing"
	ivExchangeMode    = "exchangeMode"

	ovResponse = "response"
	ovTracing  = "tracing"
)

var (
	errorDestinationIsNotAString         = errors.New("destination is not a string")
	errorDestinationTypeIsNotAString     = errors.New("destinationType is not a string")
	errorInvalidDestinationType          = errors.New("destinationType is not QUEUE or TOPIC")
	errorInvalidEmptyDestinationToSendTo = errors.New("invalid empty destination to send to")
	errorDeliveryDelayIsNotANumber       = errors.New("delivery delay is not a number")
	errorDeliveryModeIsNotAString        = errors.New("delivery mode is not a string")
	errorExpirationIsNotANumber          = errors.New("expiration is not a number")
	errorExchangeModeIsNotAString        = errors.New("exchange mode is not a string")
)

var log = logger.GetLogger("activity-tibco-ems")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	var span opentracing.Span
	if tracing := context.GetInput(ivTracing); tracing != nil {
		span = opentracing.SpanFromContext(tracing.(ctx.Context))
	}

	if span != nil {
		span = opentracing.StartSpan(
			context.TaskName(),
			opentracing.ChildOf(span.Context()))
		context.SetOutput(ovTracing, opentracing.ContextWithSpan(ctx.Background(), span))
		defer span.Finish()
	}

	setTag := func(key string, value interface{}) {
		if span != nil {
			span.SetTag(key, value)
		}
	}

	logError := func(format string, a ...interface{}) {
		str := fmt.Sprintf(format, a...)
		setTag("error", str)
		log.Error(str)
	}

	opts := ems.NewClientOptions()

	if serverURL, ok := context.GetInput(ivServerURL).(string); ok {
		setTag("serverUrl", serverURL)
		opts.SetServerUrl(serverURL)
	}
	if user, ok := context.GetInput(ivUser).(string); ok {
		opts.SetUsername(user)
	}
	if password, ok := context.GetInput(ivPassword).(string); ok {
		opts.SetPassword(password)
	} else {
		opts.SetPassword("")
	}

	client := ems.NewClient(opts)

	err = client.Connect()
	if err != nil {
		logError("Connection to EMS Server failed %v", err.Error())
	}
	defer client.Disconnect()

	content := ""
	switch v := context.GetInput(ivContent).(type) {
	case int, int64, float64, bool, json.Number:
		content = fmt.Sprintf("%v", v)
	case string:
		content = v
	default:
		var data []byte
		data, err = json.Marshal(v)
		if err != nil {
			logError("Invalid content %v", err)
			break
		}
		content = string(data)
	}
	setTag("content", content)

	destination, ok := context.GetInput(ivDestination).(string)
	if !ok {
		logError(errorDestinationIsNotAString.Error())
		return false, errorDestinationIsNotAString
	}
	if len(destination) == 0 {
		logError(errorInvalidEmptyDestinationToSendTo.Error())
		return false, errorInvalidEmptyDestinationToSendTo
	}
	setTag("destination", destination)

	destinationType, ok := context.GetInput(ivDestinationType).(string)
	if !ok {
		logError(errorDestinationTypeIsNotAString.Error())
		return false, errorDestinationTypeIsNotAString
	}
	setTag("destinationType", destinationType)

	deliveryDelay, ok := context.GetInput(ivDeliveryDelay).(int)
	if !ok {
		logError(errorDeliveryDelayIsNotANumber.Error())
		return false, errorDeliveryDelayIsNotANumber
	}

	expiration, ok := context.GetInput(ivExpiration).(int)
	if !ok {
		logError(errorExpirationIsNotANumber.Error())
		return false, errorExpirationIsNotANumber
	}

	deliveryMode, ok := context.GetInput(ivDeliveryMode).(string)
	if !ok {
		logError(errorDeliveryModeIsNotAString.Error())
		return false, errorDeliveryModeIsNotAString
	}

	exchangeMode, ok := context.GetInput(ivExchangeMode).(string)
	if !ok {
		logError(errorExchangeModeIsNotAString.Error())
		return false, errorDestinationIsNotAString
	}

	if exchangeMode == "send-only" {
		err = client.Send(destination, "QUEUE", content, deliveryDelay, deliveryMode, expiration)
		if err != nil {
			logError("Timeout occurred while trying to send to destination '%s'", destination)
			return false, err
		}
	} else {
		response, err := client.SendReceive(destination, "QUEUE", content, deliveryMode, expiration)
		if err != nil {
			logError("Timeout occurred while trying to send to destination '%s'", destination)
			return false, err
		}

		log.Debugf("Response payload: %s", response)
		context.SetOutput(ovResponse, response)

	}

	return true, nil
}
