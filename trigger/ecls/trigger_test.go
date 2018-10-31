package ecls

import (
	"context"
	"io/ioutil"

	"encoding/json"
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"testing"
)

var jsonMetadata = getJSONMetadata()

func getJSONMetadata() string {
	jsonMetadataBytes, err := ioutil.ReadFile("trigger.json")
	if err != nil {
		panic("No Json Metadata found for trigger.json path")
	}
	return string(jsonMetadataBytes)
}

// Run Once, Start Immediately
const testConfig string = `{
  "name": "ecls",
  "settings": {
		"url": "wss://logstream-api.mashery.com/ecls/subscribe/567a829c-6733-416e-86a1-f74189687708/3782cd3e-33f3-4699-930e-d48d3b2e9688?key=qUtedQriyXWfTGLBvESVCubrGtxycJRAGGl"
  },
  "handlers": [
    {
      "flowURI": "",
      "settings": {
        "handler_setting": "xxx"
      }
    }
  ]
}`

type TestRunner struct {
}

var Test action.Runner

// Run implements action.Runner.Run
func (tr *TestRunner) Run(context context.Context, action action.Action, uri string, options interface{}) (code int, data interface{}, err error) {
	log.Infof("Ran Action (Run): %v", uri)
	return 0, nil, nil
}

func (tr *TestRunner) RunAction(ctx context.Context, act action.Action, options map[string]interface{}) (results map[string]*data.Attribute, err error) {
	log.Infof("Ran Action (RunAction): %v", act)
	return nil, nil
}

func (tr *TestRunner) Execute(ctx context.Context, act action.Action, inputs map[string]*data.Attribute) (results map[string]*data.Attribute, err error) {
	log.Infof("Ran Action (Execute): %v", act)
	return nil, nil
}

func TestTrigger(t *testing.T) {

	log.Info("Testing Trigger")
	config := trigger.Config{}
	json.Unmarshal([]byte(testConfig), &config)

	factory := &wsTriggerFactory{}

	trigger := factory.New(&config)

	trigger.Start()

	defer trigger.Stop()

}
