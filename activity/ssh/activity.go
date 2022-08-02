package ssh

import (
	"bytes"
	"fmt"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
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

	input := &Input{}
	ctx.GetInputObject(input)

	hostKeyCallback, err := knownhosts.New("~/.ssh/known_hosts")
	if err != nil {
		return false, err
	}
	config := &ssh.ClientConfig{
		User: input.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(input.Password),
		},
		HostKeyCallback: hostKeyCallback,
	}
	// connect to ssh server
	conn, err := ssh.Dial("tcp", input.Host, config)
	if err != nil {
		return false, err
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return false, err

	}
	defer session.Close()

	// configure terminal mode
	modes := ssh.TerminalModes{
		ssh.ECHO: 0, // supress echo

	}
	// run terminal session
	if err := session.RequestPty("xterm", 50, 80, modes); err != nil {
		return false, err
	}
	// start remote shell
	if err := session.Shell(); err != nil {
		return false, err
	}

	var buff bytes.Buffer
	session.Stdout = &buff
	if err := session.Run(input.Command); err != nil {
		return false, err
	}
	fmt.Println(buff.String())

	output := &Output{}

	output.Result = buff.String()

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}

func init() {
	_ = activity.Register(&SSHActivity{}, New)
}

func New(ctx activity.InitContext) (activity.Activity, error) {
	return &SSHActivity{}, nil
}
