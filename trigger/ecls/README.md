# Ecls
This trigger provides your flogo application a websocket streaming connection for Mashery ECLS

## Installation

```bash
flogo install github.com/mmussett/flogo-components/trigger/ecls
```
Link for flogo web:
```
https://github.com/mmussett/flogo-components/trigger/ecls
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
      "name": "api_key",
      "type": "string"
    },
    {
      "name": "api_method_name",
      "type": "string"
    },
    {
      "name": "bytes",
      "type": "integer"
    },
    {
      "name": "cache_hit",
      "type": "string"
    },
    {
      "name": "client_transfer_time",
      "type": "number"
    },
    {
      "name": "client_transfer_time",
      "type": "number"
    },
    {
      "name": "connect_time",
      "type": "number"
    },
    {
      "name": "endpoint_name",
      "type": "string"
    },
    {
      "name": "http_method",
      "type": "string"
    },
    {
      "name": "http_status_code",
      "type": "integer"
    },
    {
      "name": "http_version",
      "type": "string"
    },
    {
      "name": "oauth_access_token",
      "type": "string"
    },
    {
      "name": "package_name",
      "type": "string"
    },
    {
      "name": "package_uuid",
      "type": "string"
    },
    {
      "name": "plan_name",
      "type": "string"
    },
    {
      "name": "plan_uuid",
      "type": "string"
    },
    {
      "name": "pre_transfer_time",
      "type": "number"
    },
    {
      "name": "qps_throttle_value",
      "type": "integer"
    },
    {
      "name": "quota_value",
      "type": "integer"
    },
    {
      "name": "referrer",
      "type": "string"
    },
    {
      "name": "remote_total_time",
      "type": "number"
    },
    {
      "name": "request_host_name",
      "type": "string"
    },
    {
      "name": "request_id",
      "type": "string"
    },
    {
      "name": "request_time",
      "type": "string"
    },
    {
      "name": "request_uuid",
      "type": "string"
    },
    {
      "name": "response_string",
      "type": "string"
    },
    {
      "name": "service_definition_endpoint_uuid",
      "type": "string"
    },
    {
      "name": "service_id",
      "type": "string"
    },
    {
      "name": "service_name",
      "type": "string"
    },
    {
      "name": "src_ip",
      "type": "string"
    },
    {
      "name": "ssl_enabled",
      "type": "string"
    },
    {
      "name": "total_request_exec_time",
      "type": "number"
    },
    {
      "name": "traffic_manager",
      "type": "string"
    },
    {
      "name": "traffic_manager_error_code",
      "type": "string"
    },
    {
      "name": "uri",
      "type": "string"
    },
    {
      "name": "user_agent",
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
"triggers": [
    {
      "id": "receive_web_socket_message",
      "ref": "github.com/mmussett/flogo-components/trigger/websocket",
      "name": "Receive WebSocket Message",
      "description": "WebSocket message handler",
      "settings": {
        "url": "wss://logstream-api.mashery.com/ecls/subscribe/567a829c-6733-416e-86a1-f74189687708/3782cd3e-33f3-4699-930e-d48d3b2e9688?key=xxx"
      },
      "handlers": [
        {
          "action": {
            "ref": "github.com/TIBCOSoftware/flogo-contrib/action/flow",
            "data": {
              "flowURI": "res://flow:web_socket_handler"
            },
            "mappings": {
              "input": [
                {
                  "mapTo": "event",
                  "type": "assign",
                  "value": "$.event"
                }
              ]
            }
          },
          "settings": {
            "handler_setting": "\"\""
          }
        }
      ]
    }
  ],
```

