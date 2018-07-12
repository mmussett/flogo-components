# WebSocket
This trigger provides your flogo application a websocket streaming connection to an endpoint

## Installation

```bash
flogo install github.com/mmussett/flogo-components/trigger/websocket
```
Link for flogo web:
```
https://github.com/mmussett/flogo-components/trigger/websocket
```

## Schema
Outputs and Endpoint:

```json
{
"settings":[
    {
      "name": "url",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "event",
      "type": "string"
    }
  ],
  "handler": {
    "settings": [
      {
        "name": "handler_setting",
        "type": "string"
      }
    ]
}
```
## Settings
| Setting   | Description    |
|:----------|:---------------|
| url  | websocket url |



## Ouputs
| Output   | Description    |
|:---------|:---------------|
| event    | The event data from the websocket server |

## Handlers
| Setting   | Description    |
|:----------|:---------------|
| N/A       | awaiting better understanding  |


## Example Configuration

Triggers are configured via the triggers.json of your application. The following is and example configuration of the WebSocket Trigger.

### Read WebSocket Data
Configure the Trigger to receive websocket events
```json
{
  "name": "websocket",
  "settings": {
		"url": "wss://localhost:4500/wsevents"
  },
  "handlers": [
    {
      "actionId": "local://testFlow2",
      "settings": {
        "handler_setting": "xxx"
      }
    }
  ]
}}
```
