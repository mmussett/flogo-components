# EMS
This trigger provides your flogo application the ability to receive JMS Text Messages from a destination

## Installation

```bash
flogo install github.com/mmussett/flogo-components/trigger/ems
```
Link for flogo web:
```
https://github.com/mmussett/flogo-components/trigger/ems
```

## Schema
Outputs and Endpoint:

```json
{
  "settings":[
    {
      "name": "serverUrl",
      "type": "string"
    },
    {
      "name": "destination",
      "type": "string"
    },
    {
      "name": "user",
      "type": "string"
    },    
    {
      "name": "password",
      "type": "string"
    }             
  ],
  "outputs": [
    {
      "name": "msgText",
      "type": "string"
    }
  ],
  "handler": {
    "settings": [
      {
        "name": "handler_setting",
        "type": "string"
      }
    ]
  }
}
```
## Settings
| Setting   | Description    |
|:----------|:---------------|
| serverUrl  | EMS Server url |
| destination | EMS Destination to receive from |
| user | EMS connection username |
| password | EMS connection password |



## Ouputs
| Output   | Description    |
|:---------|:---------------|
| msgText | Received EMS Message Text |

## Handlers
| Setting   | Description    |
|:----------|:---------------|
| N/A       | awaiting better understanding  |


## Example Configuration

Triggers are configured via the triggers.json of your application. 
The following is and example configuration of the EMS Trigger.

### Log Mashery ECLS Data
Configure the Trigger to receive EMS Messages
```json
{
  "name": "EMSSApp",
  "type": "flogo:app",
  "version": "0.0.1",
  "appModel": "1.0.0",
  "triggers": [
    {
      "id": "receive_ems_message",
      "ref": "github.com/mmussett/flogo-components/trigger/ems",
      "name": "Receive EMS Message",
      "description": "EMS message handler",
      "settings": {
        "url": "tcp://127.0.0.1:722",
        "destination": "queue.sample",
        "user": "admin",
        "password": ""
      },
      "handlers": [
        {
          "action": {
            "ref": "github.com/TIBCOSoftware/flogo-contrib/action/flow",
            "data": {
              "flowURI": "res://flow:subscriber"
            },
            "mappings": {
              "input": [
                {
                  "mapTo": "message",
                  "type": "assign",
                  "value": "$.msgText"
                }
              ]
            }
          }
        }
      ]
    }
  ]
}
```

