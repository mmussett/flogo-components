# EMS
This trigger provides your flogo application the ability to receive JMS Text Messages from a destination

## Pre-requisites

The trigger uses EMS C libraries in order to receive messages from EMS. 
Trigger has been tested against EMS 8.4 on Mac OS.

An installation of EMS 8.4 is required for this trigger to work. 

Once you have installed TIBCO EMS, you will need to make sure that the dynamic libraries are accessible.

Either copy the following dylibs to /usr/local/lib for the trigger to work:

* libtibems64.dylib
* libssl.1.0.0.dylib
* libcrypto.1.0.0.dylib

Alternatively, setting DYLD_LIBRARY_PATH or LD_LIBRARY_PATH to the location of EMS Client Libraries (/ems/8.4/lib) 
should work too (i haven't tested this).


The trigger uses ems client go package (go get github.com/mmussett/) which will need to be modified before building. 

Modify the CFLAGS and LDFLAGS paths accordingly in client.go:

```
#cgo darwin CFLAGS: -I/opt/tibco/ems/ems841/ems/8.4/include/tibems
#cgo darwin LDFLAGS: -L/opt/tibco/ems/ems841/ems/8.4/lib -ltibems64
```

## Installation

```bash
flogo install github.com/mmussett/flogo-components/trigger/ems
```
Link for flogo web:
```
https://github.com/mmussett/flogo-components/trigger/ems
```

## Schema
Outputs and Endpoint:

```json
{
  "settings":[
    {
      "name": "serverUrl",
      "type": "string"
    },
    {
      "name": "destination",
      "type": "string"
    },
    {
      "name": "user",
      "type": "string"
    },    
    {
      "name": "password",
      "type": "string"
    }             
  ],
  "outputs": [
    {
      "name": "msgText",
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
| serverUrl  | EMS Server url |
| destination | EMS Destination to receive from |
| user | EMS connection username |
| password | EMS connection password |



## Ouputs
| Output   | Description    |
|:---------|:---------------|
| msgText | Received EMS Message Text |

## Handlers
| Setting   | Description    |
|:----------|:---------------|
| N/A       | awaiting better understanding  |


## Example Configuration

Triggers are configured via the triggers.json of your application. 
The following is and example configuration of the EMS Trigger.

### Log Mashery ECLS Data
Configure the Trigger to receive EMS Messages
```json
{
  "name": "EMSSApp",
  "type": "flogo:app",
  "version": "0.0.1",
  "appModel": "1.0.0",
  "triggers": [
    {
      "id": "receive_ems_message",
      "ref": "github.com/mmussett/flogo-components/trigger/ems",
      "name": "Receive EMS Message",
      "description": "EMS message handler",
      "settings": {
        "url": "tcp://127.0.0.1:722",
        "destination": "queue.sample",
        "user": "admin",
        "password": ""
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
                  "mapTo": "message",
                  "type": "assign",
                  "value": "$.msgText"
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

## Testing

trigger_test.go is provided to test and can be invoked using:

```
$ go test -v
=== RUN   TestTrigger
2019-06-05 12:41:31.156 INFO   [trigger-ems] - Testing Trigger
2019-06-05 12:41:31.473 INFO   [trigger-ems] - event processing cycle starting
2019-06-05 12:41:35.416 INFO   [trigger-ems] - received message from EMS...
2019-06-05 12:41:35.416 INFO   [trigger-ems] - [hello, world]
2019-06-05 12:41:35.416 INFO   [trigger-ems] - event processing cycle completed
2019-06-05 12:41:35.416 INFO   [trigger-ems] - event processing cycle starting
```

Using the shipped EMS C examples you can send message via tibemsMsgProducer to fire trigger:

```
$ ./tibemsMsgProducer -queue queue.sample "hello, world"
------------------------------------------------------------------------
tibemsMsgProducer SAMPLE
------------------------------------------------------------------------
Server....................... localhost
User......................... (null)
Destination.................. queue.sample
Send Asynchronously.......... false
Message Text.................
	hello, world
------------------------------------------------------------------------
Publishing to destination 'queue.sample'
Published message: hello, world
```



