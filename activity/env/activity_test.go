package env

import (
	"io/ioutil"
	"testing"

	"fmt"
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

func TestSet(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivMethod, methodSet)
	tc.SetInput(ivEnvName, "hello")
	tc.SetInput(ivEnvValue, "world")

	_, err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set env value", err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))

}

func TestGet(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput(ivMethod, methodGet)
	tc.SetInput(ivEnvName, "GOPATH")

	_, err := act.Eval(tc)
	if err != nil {
		t.Error("unable to get env value", err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))

}
