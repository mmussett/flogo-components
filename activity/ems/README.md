# ems
This activity provides your Flogo application the ability to send EMS messages.

## Installation

```
flogo install github.com/mmussett/flogo-components/activity/ems
```

Link for flogo web:

```
https://github.com/mmussett/flogo-components/activity/ems
```


## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "content",
      "type": "string",
      "required": true
    },
    {
      "name": "destination",
      "type": "string",
      "required": true
    },
    {
      "name": "serverUrl",
      "type": "string",
      "required": true
    },
    {
      "name": "user",
      "type": "string",
      "required": true
    },
    {
      "name": "password",
      "type": "string",
      "required": false
    },
    {
      "name": "deliveryDelay",
      "type": "integer",
      "required": true
    },
    {
      "name": "deliveryMode",
      "type": "string",
      "required": true
    },
    {
      "name": "expiration",
      "type": "integer",
      "required": true
    },
    {
      "name": "tracing",
      "type": "any",
      "required": false
    }
  ],
  "outputs": [
    {
      "name": "tracing",
      "type": "any"
    }
  ]
}

```

## Inputs
| Setting     | Description    |
|:------------|:---------------|
| content     | The message to send |
| destination | The EMS queue to send the message to   |
| serverUrl   | The EMS server URL e.g. tcp://7222 |
| user        | The user name for the EMS server |
| password    | The password for the EMS server |
| tracing     | The tracing context |
| exchangeMode| Set the exchange pattern, either "send-only","send-receive","receive-only"|

## Outputs
| Setting     | Description    |
|:------------|:---------------|
| tracing     | The output tracing context |

## Configuration Example
```json
{
  "id": 2,
  "type": 1,
  "activityRef": "github.com/mmussett/flogo-contrib/activity/ems",
  "name": "ems",
  "attributes": [
    {
      "name": "content",
      "value": "test",
      "type": "string"
    },
    {
      "name": "destination",
      "value": "queue.sample",
      "type": "string"
    },
    {
      "name": "serverUrl",
      "value": "tcp://7222",
      "type": "string"
    },
    {
      "name": "user",
      "value": "admin",
      "type": "string"
    },
    {
      "name": "password",
      "value": "",
      "type": "string"
    },
    {
      "name": "deliveryDelay",
      "value": "0",
      "type": "integer"
    },
    {
      "name": "deliveryMode",
      "value": "non_persistent",
      "type": "string"
    },
    {
      "name": "expiration",
      "value": "10000",
      "type": "integer"
    }
  ],
  "inputMappings": [
    {
      "type": 1,
      "value": "${trigger.content}",
      "mapTo": "content"
    }
  ]
}
```
