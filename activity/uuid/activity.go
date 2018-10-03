package uuid

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/satori/go.uuid"
)

const (
	ivVersion   = "version"
	ivDomain    = "domain"
	ivNamespace = "namespace"
	ivName      = "name"
	ovResult    = "result"
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

	version := context.GetInput(ivVersion).(string)

	switch version {

	case "V1":
		u, err := uuid.NewV1()
		if err != nil {
			return false, err
		}
		context.SetOutput(ovResult, u.String())

	case "V2":

		domain := context.GetInput(ivDomain).(string)

		switch domain {
		case "Person":
			u, err := uuid.NewV2(uuid.DomainPerson)
			if err != nil {
				return false, err
			}
			context.SetOutput(ovResult, u.String())
		case "Group":
			u, err := uuid.NewV2(uuid.DomainGroup)
			if err != nil {
				return false, err
			}
			context.SetOutput(ovResult, u.String())
		case "Org":
			u, err := uuid.NewV2(uuid.DomainOrg)
			if err != nil {
				return false, err
			}
			context.SetOutput(ovResult, u.String())
		}

	case "V3":

		namespace := context.GetInput(ivNamespace).(string)
		name := context.GetInput(ivName).(string)
		switch namespace {
		case "DNS":
			u := uuid.NewV3(uuid.NamespaceDNS, name)
			context.SetOutput(ovResult, u.String())
		case "URL":
			u := uuid.NewV3(uuid.NamespaceURL, name)
			context.SetOutput(ovResult, u.String())
		case "OID":
			u := uuid.NewV3(uuid.NamespaceOID, name)
			context.SetOutput(ovResult, u.String())
		case "X500":
			u := uuid.NewV3(uuid.NamespaceX500, name)
			context.SetOutput(ovResult, u.String())
		}

	case "V4":

		u, err := uuid.NewV4()
		if err != nil {
			return false, err
		}
		context.SetOutput(ovResult, u.String())

	case "V5":

		namespace := context.GetInput(ivNamespace).(string)
		name := context.GetInput(ivName).(string)
		switch namespace {
		case "DNS":
			u := uuid.NewV5(uuid.NamespaceDNS, name)
			context.SetOutput(ovResult, u.String())
		case "URL":
			u := uuid.NewV5(uuid.NamespaceURL, name)
			context.SetOutput(ovResult, u.String())
		case "OID":
			u := uuid.NewV5(uuid.NamespaceOID, name)
			context.SetOutput(ovResult, u.String())
		case "X500":
			u := uuid.NewV5(uuid.NamespaceX500, name)
			context.SetOutput(ovResult, u.String())
		}

	}

	return true, nil
}
