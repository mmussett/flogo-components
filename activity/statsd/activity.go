package statsd

import (
	"errors"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/quipo/statsd"
	ctx "golang.org/x/net/context"
	"strconv"
)

const (
	ivServer     = "server"
	ivPrefix     = "prefix"
	ivMetricType = "metrictype"
	ivBucket     = "bucket"
	ivValue      = "value"
	ivTracing    = "tracing"

	ovTracing = "tracing"
)

var (
	errorSocketCreation = errors.New("failed to create socket connection")
	errorInvalidValue   = errors.New("invalid value")
)

var log = logger.GetLogger("activity-tibco-statsd")

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

	server := context.GetInput(ivServer).(string)

	prefix := context.GetInput(ivPrefix).(string)

	metricType := context.GetInput(ivMetricType).(string)

	bucket := context.GetInput(ivBucket).(string)

	statsdClient := statsd.NewStatsdClient(server, prefix)
	defer statsdClient.Close()

	err = statsdClient.CreateSocket()
	if err != nil {
		log.Error("socket creation failed %v", err.Error())
		return false, errorSocketCreation
	}


	switch metricType {
	//"absolute-int","absolute-float","decr","incr","gauge-int","gauge-float","gauge-delta-int","gauge-delta-float","timing","total"
	case "absolute-int":

		s, ok := context.GetInput(ivValue).(string)
		if !ok {
			logError(" metric requires integer value")
			return false, errorInvalidValue
		}

		value,err := strconv.ParseInt(s, 10,64)
		if err != nil {
			logError("metric requires integer value")
			return false, errorInvalidValue
		}
		statsdClient.Absolute(bucket, value)

	case "absolute-float":

		s, ok := context.GetInput(ivValue).(string)
		if !ok {
			logError("metric requires float value")
			return false, errorInvalidValue
		}

		value,err := strconv.ParseFloat(s, 64)
		if err != nil {
			logError("metric requires float value")
			return false, errorInvalidValue
		}
		statsdClient.FAbsolute(bucket, value)

	case "decr":

		s, ok := context.GetInput(ivValue).(string)
		if !ok {
			logError("metric requires integer value")
			return false, errorInvalidValue
		}

		value,err := strconv.ParseInt(s, 10,64)
		if err != nil {
			logError("metric requires integer value")
			return false, errorInvalidValue
		}
		statsdClient.Decr(bucket, value)

	case "incr":

		s, ok := context.GetInput(ivValue).(string)
		if !ok {
			logError("metric requires integer value")
			return false, errorInvalidValue
		}

		value,err := strconv.ParseInt(s, 10,64)
		if err != nil {
			logError("metric requires integer value")
			return false, errorInvalidValue
		}
		statsdClient.Incr(bucket, int64(value))

	case "gauge-int":

		s, ok := context.GetInput(ivValue).(string)
		if !ok {
			logError("metric requires integer value")
			return false, errorInvalidValue
		}

		value,err := strconv.ParseInt(s, 10,64)
		if err != nil {
			logError("metric requires integer value")
			return false, errorInvalidValue
		}

		statsdClient.Gauge(bucket, value)

	case "gauge-float":

		s, ok := context.GetInput(ivValue).(string)
		if !ok {
			logError("metric requires float value")
			return false, errorInvalidValue
		}

		value,err := strconv.ParseFloat(s, 64)
		if err != nil {
			logError("metric requires float value")
			return false, errorInvalidValue
		}
		statsdClient.FGauge(bucket, value)

	case "gauge-delta-int":

		s, ok := context.GetInput(ivValue).(string)
		if !ok {
			logError("metric requires integer value")
			return false, errorInvalidValue
		}

		value,err := strconv.ParseInt(s, 10,64)
		if err != nil {
			logError("metric requires integer value")
			return false, errorInvalidValue
		}
		statsdClient.GaugeDelta(bucket, int64(value))

	case "gauge-delta-float":

		s, ok := context.GetInput(ivValue).(string)
		if !ok {
			logError("metric requires float value")
			return false, errorInvalidValue
		}

		value,err := strconv.ParseFloat(s, 64)
		if err != nil {
			logError("metric requires float value")
			return false, errorInvalidValue
		}
		statsdClient.FGaugeDelta(bucket, value)

	case "timing":

		s, ok := context.GetInput(ivValue).(string)
		if !ok {
			logError("metric requires integer value")
			return false, errorInvalidValue
		}

		value,err := strconv.ParseInt(s, 10,64)
		if err != nil {
			logError("metric requires integer value")
			return false, errorInvalidValue
		}
		statsdClient.Timing(bucket, value)

	case "total":

		s, ok := context.GetInput(ivValue).(string)
		if !ok {
			logError("metric requires integer value")
			return false, errorInvalidValue
		}

		value,err := strconv.ParseInt(s, 10,64)
		if err != nil {
			logError("metric requires integer value")
			return false, errorInvalidValue
		}
		statsdClient.Total(bucket, value)

	default:
	}

	return true, nil
}
