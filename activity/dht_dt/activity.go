package dht_dt

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	)

const (
	ivType = "type"
	ivPin = "pin"
	ovTemp = "temp"
	ovHumidity = "humidity"
)

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

	context.SetOutput(ovTemp,0)
	context.SetOutput(ovHumidity,0)
	return true,nil

}
