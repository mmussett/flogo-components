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
      "name": "temp",
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
| pin         | True | GPIO pin to read from |
| type        | True | Sensor Type (DHT11 or DHT22) |
| boost       | False | Use boosted performance to read GPIO |

## Outputs
| Setting     | Description    |
|:------------|:---------------|
| temp | Temperature (in Celsius) value read from device |
| humidity     | Humidity value read from device |

## Configuration Example
```json
```
