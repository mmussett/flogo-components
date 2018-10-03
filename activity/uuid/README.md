# UUID
This activity provides your Flogo application the ability to generate UUID.

## Installation

```
flogo install github.com/mmussett/flogo-components/activity/uuid
```

Link for flogo web:

```
https://github.com/mmussett/flogo-components/activity/uuid
```

## UUID Versions Supported

* Version 1, based on timestamp and MAC address (RFC 4122)
* Version 2, based on timestamp, MAC address and POSIX UID/GID (DCE 1.1)
* Version 3, based on MD5 hashing (RFC 4122)
* Version 4, based on random numbers (RFC 4122)
* Version 5, based on SHA-1 hashing (RFC 4122)

## Schema
Inputs and Outputs:

```json
  "inputs": [
    {
      "name": "version",
      "type": "string",
      "allowed": ["V1", "V2","V3","V4","V5"],
      "required": true
    },
    {
      "name": "domain",
      "type": "string",
      "allowed": ["Person", "Group","Org"],
      "required": false
    },
    {
      "name": "namespace",
      "type": "string",
      "allowed": ["DNS", "URL","OID","X500"],
      "required": false
    },
    {
      "name": "name",
      "type": "string",
      "required": false
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]

```

## Inputs
| Setting     | Required | Description    |
|:------------|:---------|:---------------|
| version     | True | The UUID version (V1, V2, V3, V4, or V5) |
| domain    | False | Domain required for V2 (Person, Group, or Org)  |
| namespace     | False | Namespace required for V3 or V5 (DNS, URL, OID, or X500) |
| name     | False | Namespace name required for V3 or V5 e.g. www.tibco.com |

## Outputs
| Setting     | Description    |
|:------------|:---------------|
| result | UUID string |

## Configuration Example
```json
{
  "id": "uuid_1",
  "name": "Generate UUID",
  "description": "Activity for generating UUID",
  "activity": {
    "ref": "github.com/mmussett/flogo-components/activity/uuid",
    "input": {
      "version": "V5",
      "namespace": "DNS",
      "name": "www.tibco.com",
    }
  }
}
```
