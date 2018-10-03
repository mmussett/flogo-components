package uuid

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

func TestV1(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivVersion, "V1")

	_, err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set env value", err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))

}

func TestV2(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivVersion, "V2")
	tc.SetInput(ivDomain,"Person")

	_, err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set env value", err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))

}

func TestV3(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivVersion, "V3")
	tc.SetInput(ivNamespace,"DNS")
	tc.SetInput(ivName,"www.tibco.com")

	_, err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set env value", err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))

}

func TestV4(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivVersion, "V4")

	_, err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set env value", err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))

}

func TestV5(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivVersion, "V5")
	tc.SetInput(ivNamespace,"DNS")
	tc.SetInput(ivName,"www.tibco.com")

	_, err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set env value", err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))

}