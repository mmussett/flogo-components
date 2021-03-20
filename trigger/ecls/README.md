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
      "type": "string"
    },
    {
      "name": "cache_hit",
      "type": "string"
    },
    {
      "name": "client_transfer_time",
      "type": "string"
    },
    {
      "name": "connect_time",
      "type": "string"
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
      "type": "string"
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
      "type": "string"
    },
    {
      "name": "qps_throttle_value",
      "type": "string"
    },
    {
      "name": "quota_value",
      "type": "string"
    },
    {
      "name": "referrer",
      "type": "string"
    },
    {
      "name": "remote_total_time",
      "type": "string"
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
      "type": "string"
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
    },
    {
      "name": "log_type",
      "type": "string"
    },
    {
      "name": "ingestion_time",
      "type": "string"
    },
        {
      "name": "org_uuid",
      "type": "string",
      "description": ""
    },
    {
      "name": "org_name",
      "type": "string",
      "description": ""
    },
    {
      "name": "sub_org_uuid",
      "type": "string",
      "description": ""
    },
    {
      "name": "sub_org_name",
      "type": "string",
      "description": ""
    },
    {
      "name": "asCSV",
      "type": "string"
    },
    {
      "name": "asObject",
      "type": "object"
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
| api_key | API key used by application. Provided both when used by itself and in relation to an OAuth2 based call |
| api_method_name | Name of Method as configured via the Method detection setting on the endpoint |
| bytes | Bytes in response |
| cache_hit | 1=True / 0=False on cache hit served by TIBCO Mashery |
| client_transfer_time | Total time transferring from Mashery to client on outbound |
| connect_time | Total time Traffic Manager negotiating connection with customer origin servers(0 value means reusing existing connection) |
| endpoint_name | Name of endpoint |
| http_method | HTTP method (get, post, etc.) |
| http_status_code | HTTP status code |
| http_version | HTTP version |
| oauth_access_token | OAuth access token value |
| package_name | Name of package |
| package_uuid | Unique ID for the package, used with V3 API |
| plan_name | Name of plan |
| plan_uuid | Unique GUID for plan, used with v3 API |
| pre_transfer_time | Process time by Mashery prior to transfer |
| qps_throttle_value | Queries per second throttle count value |
| quota_value | Count against call quota |
| referrer | Client referrer information |
| remote_total_time | Total call time spent waiting for response from origin by Mashery |
| request_host_name | Name of host invoked by call |
| request_id | ID of request (time(epoch)+serviceid+API_key) |
| request_time | Date/time of request in ISO 8601 |
| request_uuid | Unique id for request |
| response_string | Response string which includes source of error (Mashery vs origin) |
| service_definition_endpoint_uuid | UUID for endpoint, used with v3 API |
| service_id | Service ID (Also sometimes known as SPKEY or Servicekey) |
| service_name | Name of service in which endpoint is located |
| src_ip | Source IP address of Client |
| ssl_enabled | 1=True/ 0=False SSL used in inbound connection|
| total_request_exec_time | Total time from receipt of request to completion of response to client, formerly exec_time |
| traffic_manager | Traffic manager host name |
| traffic_manager_error_code | Error code returned by Traffic Manager |
| uri | URI (with max length imposed) |
| user_agent | User Agent of client |
| log_type | |
| ingestion_time | |
| org_uuid | Unique identifier for a Parent organization |
| org_name | Organization Name for the Parent organization |
| sub_org_uuid | Unique identifier for the Sub organization |
| sub_org_name | Sub Organization Name |
 
| asCSV | |
| asObject | |

## Handlers
| Setting   | Description    |
|:----------|:---------------|
| N/A       | awaiting better understanding  |


## Example Configuration

Triggers are configured via the triggers.json of your application. The following is and example configuration of the Mashery ECLS Trigger.

### Log Mashery ECLS Data
Configure the Trigger to receive Mashery ECLS events
```json
{
  "name": "MasheryECLSApp",
  "type": "flogo:app",
  "version": "0.0.1",
  "appModel": "1.0.0",
  "triggers": [
    {
      "id": "receive_mashery_ecls_message",
      "ref": "github.com/mmussett/flogo-components/trigger/ecls",
      "name": "Receive Mashery ECLS Message",
      "description": "ECLS message handler",
      "settings": {
        "url": "wss://logstream-api.mashery.com/ecls/subscribe/567a829c-6733-416e-86a1-f74189687708/3782cd3e-33f3-4699-930e-d48d3b2e9688?key=xyz"
      },
      "handlers": [
        {
          "action": {
            "ref": "github.com/TIBCOSoftware/flogo-contrib/action/flow",
            "data": {
              "flowURI": "res://flow:subscriber"
            },
            "mappings": {
              "input": [
                {
                  "mapTo": "url",
                  "type": "assign",
                  "value": "$.uri"
                }
              ]
            }
          }
        }
      ]
    }
  ]
}
```

