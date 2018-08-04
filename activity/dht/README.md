# DHT
This activity provides your Flogo application the ability to read a DHT11 or 22 sensor in a Raspberry PI environment.
## Installation

```
flogo install github.com/mmussett/flogo-components/activity/dht
```

Link for flogo web:

```
https://github.com/mmussett/flogo-components/activity/dht
```


## Schema
Inputs and Outputs:

```json
{
  "inputs": [
    {
      "name": "pin",
      "type": "integer",
      "required": true
    },
    {
      "name": "type",
      "allowed": [
        "DHT22",
        "DHT11"
      ],
      "type": "string",
      "required": true,
    },
    {
      "name": "boost",
      "type": "boolean"
    }
  ],
  "outputs": [
    {
      "name": "temperature",
      "type": "number"
    },
    {
      "name": "humidity",
      "type": "number"
    }
  ]
}

```

## Inputs
| Setting     | Required | Description    |
|:------------|:---------|:---------------|
| method     | True | The method type (GET, SET) |
| envName    | True | The environment variable name |
| envValue     | False | The value to set the environment variable to |

## Outputs
| Setting     | Description    |
|:------------|:---------------|
| result |
| tracing     | The output tracing context |

## Configuration Example
```json
```
