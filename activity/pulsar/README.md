# pulsar
This activity provides your Flogo application the ability to send Apache Pulsar messages.

## Installation

```
flogo install github.com/mmussett/flogo-components/activity/pulsar
```

Link for flogo web:

```
https://github.com/mmussett/flogo-components/activity/pulsar
```


## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "url",
      "type": "string",
      "required": true
    },
    {
      "name": "topic",
      "type": "string",
      "required": true
    },
    {
      "name": "sendTimeout",
      "type": "integer",
      "required": true
    },
    {
      "name": "payload",
      "type": "string",
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
      "name": "response",
      "type": "string"
    },
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
| url   | The Pulsar server URL e.g. pulsar://localhost:6650 |
| topic     | The Pulsar topic to publish the payload to |
| sendTimeout | The time (in seconds) to block waiting on acknowledgement from Broker |
| payload | The payload for the message  |


## Outputs
| Setting     | Description    |
|:------------|:---------------|
| tracing     | The output tracing context |

## Configuration Example
```json
{
  "id": 2,
  "type": 1,
  "activityRef": "github.com/mmussett/flogo-contrib/activity/pulsar",
  "name": "ems",
  "attributes": [
    {
      "name": "url",
      "value": "pulsar://localhost:6650",
      "type": "string"
    },
    {
      "name": "topic",
      "value": "topic.sample",
      "type": "string"
    },
    {
      "name": "sendTimeout",
      "type": "integer",
      "required": true
    },
    {
      "name": "payload",
      "value": "Hello, World",
      "type": "string"
    },
    {
      "name": "tracing",
      "value": false,
      "type": "any"
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
