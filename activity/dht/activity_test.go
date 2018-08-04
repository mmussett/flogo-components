package dht

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
	tc.SetInput(ivPin, 18)
	tc.SetInput(ivType, "DHT22")
	tc.SetInput(ivBoost, true)

	_, err := act.Eval(tc)
	if err != nil {
		t.Error("unable to read sensor", err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovTemp))
	fmt.Println(tc.GetOutput(ovHumidity))

}
