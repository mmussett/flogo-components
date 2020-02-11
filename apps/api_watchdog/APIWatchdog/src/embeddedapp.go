// Do not change this file, it has been generated using flogo-cli
// If you change it and rebuild the application your changes might get lost
package main

// embedded flogo app descriptor file
const flogoJSON string = `{
  "name": "APIWatchdog",
  "type": "flogo:app",
  "version": "0.0.1",
  "description": "",
  "appModel": "1.1.0",
  "imports": [
    "github.com/project-flogo/contrib/function/coerce",
    "github.com/project-flogo/contrib/trigger/timer",
    "github.com/project-flogo/flow",
    "github.com/project-flogo/legacybridge",
    "github.com/mmussett/flogo-components/activity/statsd",
    "github.com/project-flogo/contrib/activity/rest"
  ],
  "triggers": [
    {
      "id": "timer",
      "ref": "#timer",
      "settings": null,
      "handlers": [
        {
          "settings": {
            "repeatInterval": "1m"
          },
          "actions": [
            {
              "ref": "#flow",
              "settings": {
                "flowURI": "res://flow:apiname"
              }
            }
          ]
        }
      ]
    }
  ],
  "resources": [
    {
      "id": "flow:apiname",
      "data": {
        "name": "__APINAME__",
        "tasks": [
          {
            "id": "rest_2",
            "name": "Invoke API",
            "description": "Invokes a REST Service",
            "activity": {
              "ref": "#rest",
              "settings": {
                "method": "GET",
                "uri": "https://api.openweathermap.org/data/2.5/weather?q=London\u0026appid=0e5cf0210c36ca98cf915df7f355744d"
              }
            }
          },
          {
            "id": "statsd_3",
            "name": "Publish to Statsd",
            "description": "Publish metrics to statsd",
            "activity": {
              "ref": "#statsd",
              "input": {
                "bucket": "__APINAME__",
                "prefix": "flogo.api.statsd.",
                "server": "192.168.64.3:31091",
                "metrictype": "absolute-int",
                "value": "=coerce.toString($activity[rest_2].status)"
              }
            }
          }
        ],
        "links": [
          {
            "from": "rest_2",
            "to": "statsd_3"
          }
        ]
      }
    }
  ]
}`
const engineJSON string = ``

func init () {
	cfgJson = flogoJSON
	cfgEngine = engineJSON
}
