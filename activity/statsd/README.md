# statsd
This activity provides your Flogo application the ability to send statsd events.

## Installation

```
flogo install github.com/mmussett/flogo-components/activity/statsd
```

Link for flogo web:

```
https://github.com/mmussett/flogo-components/activity/statsd
```

## Schema
Inputs and Outputs:

```json
"inputs":[
    {
      "name": "server",
      "type": "string",
      "required": true
    },
    {
      "name": "prefix",
      "type": "string",
      "required": true
    },
    {
      "name": "metrictype",
      "type": "string",
      "required": true,
      "allowed": ["absolute-int","absolute-float","decr","incr","gauge-int","gauge-float","gauge-delta-int","gauge-delta-float","timing","total"]
    },
    {
      "name": "bucket",
      "type": "string",
      "required": true
    },
    {
      "name": "value",
      "type": "any",
      "required": true
    },
    {
      "name": "tracing",
      "type": "any",
      "required": false
    }
  ]
```

## Inputs
| Setting     | Description    |
|:------------|:---------------|
| server      | The statsd server to send to e.g. 127.0.0.1:8125 |
| prefix      | The statsd prefix value e.g. flogo.stats. |
| metrictype   | The statsd metric type to send (absolute-int,absolute-float,decr,incr,gauge-int,gauge-float,gauge-delta-int,gauge-delta-float,timing,total) |
| bucket        | The statsd bucket name |
| value    | The statsd value to record |
| tracing     | The tracing context |

## Outputs
| Setting     | Description    |
|:------------|:---------------|
| tracing     | The output tracing context |

## Configuration Example
```json
{
  "id": 2,
  "type": 1,
  "activityRef": "github.com/mmussett/flogo-contrib/activity/statsd",
  "name": "statsd",
  "attributes": [
    {
      "name": "server",
      "value": "0.0.0.0:8125",
      "type": "string"
    },
    {
      "name": "prefix",
      "value": "flogo.stats.",
      "type": "string"
    },
    {
      "name": "metrictype",
      "value": "absolute-int",
      "type": "string"
    },
    {
      "name": "bucket",
      "value": "bytes-sent",
      "type": "string"
    },
    {
      "name": "value",
      "value": 534,
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
