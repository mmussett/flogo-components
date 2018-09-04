package amqp

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

const (
	uri = "amqp://guest:guest@localhost:5672/"
	exchangeName = "amqp.direct"
	exchangeType = "direct"
	routingKey = "test-key"
	body = "hello,world"
	reliable = true
)
func TestShift(t *testing.T) {


	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput(ivUri,uri)
	tc.SetInput(ivExchangeName, exchangeName)
	tc.SetInput(ivExchangeType , exchangeType)
	tc.SetInput(ivRoutingKey, routingKey)
	tc.SetInput(ivBody, body)
	tc.SetInput(ivReliable, reliable)



	_, err := act.Eval(tc)
	if err != nil {
		t.Error("unable to eval", err)
		t.Fail()
	}


	fmt.Println(tc.GetOutput(ovResult).(string))

}
