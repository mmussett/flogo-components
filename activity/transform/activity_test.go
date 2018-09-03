package transform

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


func Test1(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivInput, `{"input":"input value"}`)
	tc.SetInput(ivSpec, `[{"operation": "shift", "spec": {"output": "input"}}]`)


	_, err := act.Eval(tc)
	if err != nil {
		t.Error("unable to eval", err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovOutput))

}
