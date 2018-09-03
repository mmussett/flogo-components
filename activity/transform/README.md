# Transform
This activity provides your Flogo application the ability to run transformation on Json input strings using Kazaam transformations.

## Installation

```
flogo install github.com/mmussett/flogo-components/activity/transform
```

Link for flogo web:

```
https://github.com/mmussett/flogo-components/activity/transform
```


## Schema
Inputs and Outputs:

```json
{
  "inputs": [
    {
      "name": "input",
      "type": "string",
      "required": true
    },
    {
      "name": "spec",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "output",
      "type": "string"
    }
  ]
}

```

## Inputs
| Setting     | Required | Description    |
|:------------|:---------|:---------------|
| input     | True | Input JSON string to transform |
| spec | True | Kazaam transformation specification   |

## Outputs
| Setting     | Description    |
|:------------|:---------------|
| output | Transformed JSON string

## Configuration Example
```json
```
