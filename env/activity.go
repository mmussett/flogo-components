package env

import (
	"os"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

const (

	methodGet = "GET"
	methodSet = "SET"

	ivMethod     = "method"
	ivEnvName    = "envName"
	ivEnvValue   = "envValue"

	ovResult  = "result"

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


	method := context.GetInput(ivMethod).(string)
	envName := context.GetInput(ivEnvName).(string)
	envValue := context.GetInput(ivEnvValue).(string)


	switch method {

	case methodGet:

		// Set the output value in the context
		context.SetOutput(ovResult, os.Getenv(envName))
		return true, nil

	case methodSet:

		os.Setenv(envName, envValue)
		context.SetOutput(ovResult, "OK")

		return true, nil

	}

	// Set the output value in the context
	context.SetOutput(ovResult, "NOK")
	return true, nil
}
