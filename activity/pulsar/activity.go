package pulsar

import (
	ctx "context"
	"errors"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/opentracing/opentracing-go"
)

const (
	ivUrl     = "url"
	ivTopic   = "topic"
	ivPayload = "payload"
	ivTracing = "tracing"

	ovResponse = "response"
	ovTracing  = "tracing"
)

var (
	errorUrlIsNotAString = errors.New("url is not a string")
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

	url, ok := context.GetInput(ivUrl).(string)

	if !ok {
		logError(errorUrlIsNotAString.Error())
		return false, errorUrlIsNotAString
	}

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: url,
	})

	defer client.Close()

	topic, _ := context.GetInput(ivTopic).(string)

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,

	})

	payload, _ := context.GetInput(ivPayload).(string)

	ctx := ctx.Background()


	err = producer.Send(ctx, &pulsar.ProducerMessage{
		Payload: []byte(payload),
	})

	if err != nil {
		context.SetOutput(ovResponse, "false")
	} else {
		context.SetOutput(ovResponse, "true")
	}
	defer producer.Close()

	return true, nil
}
