// Do not change this file, it has been generated using flogo-cli
// If you change it and rebuild the application your changes might get lost
package main

// embedded flogo app descriptor file
const flogoJSON string = `{
  "name": "ems_trigger_app",
  "type": "flogo:app",
  "version": "0.0.1",
  "description": "My flogo application description",
  "appModel": "1.1.0",
  "imports": [
    "github.com/project-flogo/contrib/trigger/rest",
    "github.com/project-flogo/flow",
    "github.com/mmussett/flogo-components/trigger/ems",
    "github.com/project-flogo/legacybridge",
    "github.com/project-flogo/contrib/activity/log"
  ],
  "triggers": [
    {
      "id": "receive_ems_trigger",
      "ref": "github.com/mmussett/flogo-components/trigger/ems",
      "settings": {
        "destination": "topic.sample",
        "destinationType": "topic",
        "password": "",
        "serverURL": "tcp://127.0.0.1:7222",
        "user": "admin"
      },
      "handlers": [
        {
          "settings": null,
          "actions": [
            {
              "ref": "#flow",
              "settings": {
                "flowURI": "res://flow:simple_flow"
              },
              "input": {
                "in": "=$.data"
              }
            }
          ]
        }
      ]
    }
  ],
  "resources": [
    {
      "id": "flow:simple_flow",
      "data": {
        "name": "simple_flow",
        "metadata": {
          "input": [
            {
              "name": "in",
              "type": "string",
              "value": "test"
            }
          ],
          "output": [
            {
              "name": "out",
              "type": "string"
            }
          ]
        },
        "tasks": [
          {
            "id": "log",
            "name": "Log Message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=$flow.in",
                "addDetails": "false"
              }
            }
          }
        ],
        "links": []
      }
    }
  ]
}`

func init () {
	cfgJson = flogoJSON
}
