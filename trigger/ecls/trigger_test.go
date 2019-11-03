package ecls

import (
	"encoding/json"
	"github.com/project-flogo/core/support"
	"testing"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/support/test"
	"github.com/project-flogo/core/trigger"
	"github.com/stretchr/testify/assert"
)

const testConfig string = `{
	"id": "flogo-ecls-trigger",
	"ref": "github.com/mmussett/flogo-components/trigger/ecls",
	"settings": {
      "url": "wss://logstream-api.mashery.com/ecls/subscribe/c7e8e2d5-ff91-42eb-9885-10f2aa2cc3f5/b6d350b1-fb23-4059-a8ee-578c88531fc8?key=bBhzYwNuMRKrMHDqmjmtsqRFKFpCvNhmuue"
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
