package statsd

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
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

func TestEval_AbsoluteInt(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput(ivServer, "0.0.0.0:8125")
	tc.SetInput(ivMetricType, "absolute-int")
	tc.SetInput(ivValue, 10)
	tc.SetInput(ivPrefix, "flogo-statsd.")
	tc.SetInput(ivBucket, "my.metric")
	//setup attrs

	act.Eval(tc)

	//check result attr
}

func TestEval_AbsoluteFloat(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput(ivServer, "0.0.0.0:8125")
	tc.SetInput(ivMetricType, "absolute-float")
	tc.SetInput(ivValue, 10.1234567890)
	tc.SetInput(ivPrefix, "flogo-statsd.")
	tc.SetInput(ivBucket, "my.metric")
	//setup attrs

	act.Eval(tc)

	//check result attr
}

func TestEval_Decr(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput(ivServer, "0.0.0.0:8125")
	tc.SetInput(ivMetricType, "decr")
	tc.SetInput(ivValue, 1)
	tc.SetInput(ivPrefix, "flogo-statsd.")
	tc.SetInput(ivBucket, "my.metric")
	//setup attrs

	act.Eval(tc)

	//check result attr
}

func TestEval_Incr(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput(ivServer, "0.0.0.0:8125")
	tc.SetInput(ivMetricType, "incr")
	tc.SetInput(ivValue, 1)
	tc.SetInput(ivPrefix, "flogo-statsd.")
	tc.SetInput(ivBucket, "my.metric")
	//setup attrs

	act.Eval(tc)

	//check result attr
}

func TestEval_GaugeInt(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput(ivServer, "0.0.0.0:8125")
	tc.SetInput(ivMetricType, "gauge-int")
	tc.SetInput(ivValue, 1)
	tc.SetInput(ivPrefix, "flogo-statsd.")
	tc.SetInput(ivBucket, "my.metric")
	//setup attrs

	act.Eval(tc)

	//check result attr
}

func TestEval_GaugeFloat(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput(ivServer, "0.0.0.0:8125")
	tc.SetInput(ivMetricType, "gauge-float")
	tc.SetInput(ivValue, 1.23456)
	tc.SetInput(ivPrefix, "flogo-statsd.")
	tc.SetInput(ivBucket, "my.metric")
	//setup attrs

	act.Eval(tc)

	//check result attr
}

func TestEval_GaugeDeltaInt(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput(ivServer, "0.0.0.0:8125")
	tc.SetInput(ivMetricType, "gauge-delta-int")
	tc.SetInput(ivValue, 1)
	tc.SetInput(ivPrefix, "flogo-statsd.")
	tc.SetInput(ivBucket, "my.metric")
	//setup attrs

	act.Eval(tc)

	//check result attr
}

func TestEval_GaugeDeltaFloat(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput(ivServer, "0.0.0.0:8125")
	tc.SetInput(ivMetricType, "gauge-delta-float")
	tc.SetInput(ivValue, 1.2345)
	tc.SetInput(ivPrefix, "flogo-statsd.")
	tc.SetInput(ivBucket, "my.metric")
	//setup attrs

	act.Eval(tc)

	//check result attr
}

func TestEval_Total(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput(ivServer, "0.0.0.0:8125")
	tc.SetInput(ivMetricType, "total")
	tc.SetInput(ivValue, 1)
	tc.SetInput(ivPrefix, "flogo-statsd.")
	tc.SetInput(ivBucket, "my.metric")
	//setup attrs

	act.Eval(tc)

	//check result attr
}
