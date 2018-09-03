package transform

import (
	"io/ioutil"
	"strings"
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

const testJSONInput = `{"rating":{"example":{"value":3},"primary":{"value":3}}}`

func TestShift(t *testing.T) {

	spec := `[{"operation": "shift", "spec": {"Rating": "rating.primary.value", "example.old": "rating.example"}}]`
	jsonOut :=  `{"Rating":3,"example":{"old":{"value":3}}}`

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivContent, testJSONInput)
	tc.SetInput(ivSpec, spec)


	_, err := act.Eval(tc)
	if err != nil {
		t.Error("unable to eval", err)
		t.Fail()
	}

	if strings.Compare(tc.GetOutput(ovResult).(string),jsonOut) == 1 {
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult).(string))

}
