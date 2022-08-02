package ssh

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
)

type SSHActivity struct {
}

var activityLog = log.ChildLogger(log.RootLogger(), "tibco-activity-ssh")

var activityMd = activity.ToMetadata(&Input{}, &Output{})

func (S SSHActivity) Metadata() *activity.Metadata {
	return activityMd
}

func (S SSHActivity) Eval(ctx activity.Context) (done bool, err error) {
	activityLog.Info("Executing SSH activity")

	return true, nil
}

func init() {
	_ = activity.Register(&SSHActivity{}, New)
}

func New(ctx activity.InitContext) (activity.Activity, error) {
	return &SSHActivity{}, nil
}
