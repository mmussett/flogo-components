package ems

import (
	"context"
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/mmussett/ems"
	"os"
	"os/signal"
)

// log is the default package logger
var log = logger.GetLogger("trigger-ems")

// udpTriggerFactory My Trigger factory
type emsTriggerFactory struct {
	metadata *trigger.Metadata
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &emsTriggerFactory{metadata: md}
}

//New Creates a new trigger instance for a given id
func (t *emsTriggerFactory) New(config *trigger.Config) trigger.Trigger {
	return &emsTrigger{metadata: t.metadata, config: config}
}

// emsrigger is a stub for your Trigger implementation
type emsTrigger struct {
	metadata *trigger.Metadata
	runner   action.Runner
	config   *trigger.Config
	handlers []*trigger.Handler
}

// Metadata implements trigger.Trigger.Metadata
func (t *emsTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

func (t *emsTrigger) Initialize(ctx trigger.InitContext) error {
	log.Debug("Initialize")
	t.handlers = ctx.GetHandlers()
	return nil

}

// Start implements trigger.Trigger.Start
func (t *emsTrigger) Start() error {

	// start the trigger
	log.Debug("Start")

	// Get parms
	urlParam := t.config.GetSetting("serverUrl")
	usernameParam := t.config.GetSetting("user")
	passwordParam := t.config.GetSetting("password")
	destinationParam := t.config.GetSetting("destination")

	ops := ems.NewClientOptions().SetServerUrl(urlParam).SetUsername(usernameParam).SetPassword(passwordParam)

	c := ems.NewClient(ops).(*ems.Client)

	err := c.Connect()
	if err != nil {
		log.Error(err)
		return err
	}

	log.Debugf("EMS connection established to: [%s]", urlParam)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		defer c.Disconnect()
		for {
			log.Debug("event processing cycle starting")
			msgText, err := c.Receive(destinationParam)

			if err != nil {
				log.Errorf("read:", err)
				return
			}

			log.Debug("received message from EMS")

			trgData := make(map[string]interface{})
			trgData["msgText"] = msgText

			for _, handler := range t.handlers {
				results, err := handler.Handle(context.Background(), trgData)
				if err != nil {
					log.Error("Error starting action: %v", err)
				}
				log.Debugf("Ran Handler: [%s]", handler)
				log.Debugf("Results [%v]", results)
			}

			log.Debug("event processing cycle completed")
		}
	}()

	for {
		select {
		case <-interrupt:
			log.Debug("interrupt")
			c.Disconnect()
			return nil
		}
	}

	return nil
}

// Stop implements trigger.Trigger.Start
func (t *emsTrigger) Stop() error {
	// stop the trigger
	return nil
}
