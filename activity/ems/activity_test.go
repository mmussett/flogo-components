package ems

import (
	"context"
	"io/ioutil"
	"net"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	opentracing "github.com/opentracing/opentracing-go"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestReceiveOnly(t *testing.T) {

	_, err := net.Dial("tcp", "127.0.0.1:7222")
	if err != nil {
		t.Log("EMS Server is not available, skipping test...")
		return
	}

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivContent, `{"test": "hello world"}`)
	tc.SetInput(ivDestination, "queue.sample")
	tc.SetInput(ivServerURL, "tcp://127.0.0.1:7222")
	tc.SetInput(ivUser, "admin")
	tc.SetInput(ivPassword, "")
	tc.SetInput(ivDestinationType, "queue")
	tc.SetInput(ivDeliveryDelay, 0)
	tc.SetInput(ivDeliveryMode, "non_persistent")
	tc.SetInput(ivExpiration, 10000)
	tc.SetInput(ivTimeout,10000)
	tc.SetInput(ivExchangeMode, "receive-only")

	span := opentracing.StartSpan("test")
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	tc.SetInput(ivTracing, ctx)
	defer span.Finish()

	act.Eval(tc)

	msg := tc.GetOutput(ovResponse)
	t.Log(msg)
	//check result attr
	tracing := tc.GetOutput(ovTracing)
	if tracing == nil {
		t.Error("tracing is nil")
	}

}

func TestSendOnly(t *testing.T) {
	_, err := net.Dial("tcp", "127.0.0.1:7222")
	if err != nil {
		t.Log("EMS Server is not available, skipping test...")
		return
	}

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivContent, `{"test": "hello world"}`)
	tc.SetInput(ivDestination, "queue.sample")
	tc.SetInput(ivServerURL, "tcp://127.0.0.1:7222")
	tc.SetInput(ivUser, "admin")
	tc.SetInput(ivPassword, "")
	tc.SetInput(ivDestinationType, "queue")
	tc.SetInput(ivDeliveryDelay, 0)
	tc.SetInput(ivDeliveryMode, "non_persistent")
	tc.SetInput(ivExpiration, 1000000)
	tc.SetInput(ivTimeout,10000)
	tc.SetInput(ivExchangeMode, "send-only")

	span := opentracing.StartSpan("test")
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	tc.SetInput(ivTracing, ctx)
	defer span.Finish()

	act.Eval(tc)

	//check result attr
	tracing := tc.GetOutput(ovTracing)
	if tracing == nil {
		t.Error("tracing is nil")
	}

}

func TestSendReceive(t *testing.T) {
	_, err := net.Dial("tcp", "127.0.0.1:7222")
	if err != nil {
		t.Log("EMS Server is not available, skipping test...")
		return
	}

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivContent, `{"test": "hello world"}`)
	tc.SetInput(ivDestination, "queue.sample")
	tc.SetInput(ivServerURL, "tcp://127.0.0.1:7222")
	tc.SetInput(ivUser, "admin")
	tc.SetInput(ivPassword, "")
	tc.SetInput(ivDestinationType, "queue")
	tc.SetInput(ivDeliveryDelay, 0)
	tc.SetInput(ivDeliveryMode, "non_persistent")
	tc.SetInput(ivExpiration, 10000)
	tc.SetInput(ivTimeout,10000)
	tc.SetInput(ivExchangeMode, "send-receive")

	span := opentracing.StartSpan("test")
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	tc.SetInput(ivTracing, ctx)
	defer span.Finish()

	act.Eval(tc)

	msg := tc.GetOutput(ovResponse)
	t.Log(msg)
	//check result attr
	tracing := tc.GetOutput(ovTracing)
	if tracing == nil {
		t.Error("tracing is nil")
	}

	response := tc.GetOutput(ovResponse)
	if response == nil {
		t.Error("response is nil")
	}
}
