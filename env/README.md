# Env
This activity provides your Flogo application the ability to get/set environment variables.

## Installation

```
flogo install github.com/mmussett/flogo-components/activity/env
```

Link for flogo web:

```
https://github.com/mmussett/flogo-components/activity/env
```


## Schema
Inputs and Outputs:

```json
{
  "inputs": [
    {
      "name": "method",
      "type": "string",
      "allowed": [
        "GET",
        "SET"
      ],
      "required": true
    },
    {
      "name": "envName",
      "type": "string"
    },
    {
      "name": "envValue",
      "type": "string"
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
