package ems

import (
	"encoding/json"
	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/support/test"
	"testing"

	"github.com/project-flogo/core/support"
	"github.com/project-flogo/core/trigger"
	"github.com/stretchr/testify/assert"
)

const testConfig string = `{
	"id": "flogo-ems-trigger",
	"ref": "github.com/mmussett/flogo-components/trigger/ems",
	"settings": {
      "serverURL": "tcp://localhost:7222",
      "destination":"queue.sample",
      "destinationType":"queue",
      "timeout":"0",
      "username":"admin",
      "password":""
	},
	"handlers": [
	  {
			"action":{
				"id":"dummy"
			},
			"settings": {
		  	
			}
	  }
	]
	
  }`

func TestTrigger_Register(t *testing.T) {

	ref := support.GetRef(&Trigger{})
	f := trigger.GetFactory(ref)
	assert.NotNil(t, f)
}

func TestEclsTrigger_Initialize(t *testing.T) {
	f := &Factory{}

	config := &trigger.Config{}
	err := json.Unmarshal([]byte(testConfig), config)
	assert.Nil(t, err)

	actions := map[string]action.Action{"dummy": test.NewDummyAction(func() {
		//do nothing
	})}

	trg, err := test.InitTrigger(f, config, actions)
	assert.Nil(t, err)
	assert.NotNil(t, trg)

	err = trg.Start()
	assert.Nil(t, err)
	err = trg.Stop()
	assert.Nil(t, err)

}
