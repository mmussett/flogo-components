# AMQP Publisher
This activity provides your Flogo application the ability to publish to AMQP.

## Installation

```
flogo install github.com/mmussett/flogo-components/activity/amqp
```

Link for flogo web:

```
https://github.com/mmussett/flogo-components/activity/amqp
```


## Schema
Inputs and Outputs:

```json
{
  "inputs": [
    {
      "name": "uri",
      "type": "string",
      "required": true
    },
    {
      "name": "exchangeName",
      "type": "string",
      "required": true
    },
    {
      "name": "exchangeType",
      "type": "string",
      "required": true
    },
    {
      "name": "routingKey",
      "type": "string",
      "required": true
    },
    {
      "name": "body",
      "type": "any",
      "required": true
    },
    {
      "name": "reliable",
      "type": "boolean",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}

```

## Inputs
| Setting     | Required | Description    |
|:------------|:---------|:---------------|
| uri     | True | AMQP URI |
| exchangeName | True | Durable AMQP exchange name   |
| exchangeType | True | Exchange Type (direct, fanout, topic, x-custom)   |
| routingKey | True | AMQP routing key   |
| body | True | Body of message   |
| reliable | True | Wait for the publisher confirmation before completing   |

## Outputs
| Setting     | Description    |
|:------------|:---------------|
| result | Returns 'OK' |

## Configuration Example
```json
{
"id": "amqp-publish_1",
"name": "AMQP Publish",
"description": "AMQP publisher activity.",
    "activity": {
        "ref": "github.com/mmussett/flogo-components/activity/amqp-publish",
        "input": {
            "uri": "amqp://guest:guest@localhost:5672/",
            "exchangeName": "amqp.direct",
            "exchangeType": "direct",
            "routingKey": "test-key",
            "body": "Hello, World",
            "reliable": false
        }
    }
}
```

