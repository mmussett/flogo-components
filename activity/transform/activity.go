package transform

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/qntfy/kazaam"
)

const (

	ivInput      = "input"
	ivSpec       = "spec"
	ovOutput     = "output"
)

var log = logger.GetLogger("activity-tibco-transform")

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


	input := context.GetInput(ivInput).(string)
	spec := context.GetInput(ivSpec).(string)

	k, _ := kazaam.NewKazaam(spec)

	output, _ := k.TransformJSONStringToString(input)

	log.Debugf("Response payload: %s", output)

	// Set the output value in the context
	context.SetOutput(ovOutput, output)
	return true, nil
}
