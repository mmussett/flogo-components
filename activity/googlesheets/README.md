# Google Sheets
This activity provides your Flogo application the ability to manipulate Google Sheets.

## Installation

```
flogo install github.com/mmussett/flogo-components/activity/googlesheets
```

Link for flogo web:

```
https://github.com/mmussett/flogo-components/activity/googlesheets
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
        "PING",
        "GET",
        "SET",
        "DELETE",
        "EXISTS",
        "EXPIRE",
        "PERSIST",
        "PEXPIRE",
        "TTL",
        "PTTL",
        "RENAME",
        "APPEND",
        "FLUSHDB",
        "FLUSHALL",
        "DECR",
        "INCR"
      ],
      "required": true
    },
    {
      "name": "address",
      "type": "string",
      "required": true
    },
    {
      "name": "password",
      "type": "string"
    },
    {
      "name": "database",
      "type": "integer",
      "required": true
    },
    {
      "name": "key",
      "type": "string"
    },
    {
      "name": "value",
      "type": "string"
    },
    {
      "name": "expiration",
      "type": "integer"
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
| method     | True | The method type (PING, GET, SET, DELETE) |
| address | True | Address of Redis server to connect to e.g. localhost:12000   |
| password | False |  The AUTH password required to connect to Redis |
| database | True | The Redis database number to connect to e.g. 0 |
| keyName    | False | The key (used for get, set, or delete action) |
| keyValue     | False | The value of the key (used for set action) |
| expiration | False | The expiration value for key-value pair (set to -1 for persistent key-value) |

## Outputs
| Setting     | Description    |
|:------------|:---------------|
| result |
| tracing     | The output tracing context |

## Configuration Example
```json
```
