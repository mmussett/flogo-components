package websocket

import (
	"context"
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	"net/url"
	"crypto/tls"
	"github.com/gorilla/websocket"

	"strings"
	"time"
	"os/signal"
	"os"
)

// log is the default package logger
var log = logger.GetLogger("trigger-websocket")

// udpTriggerFactory My Trigger factory
type wsTriggerFactory struct {
	metadata *trigger.Metadata
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &wsTriggerFactory{metadata: md}
}

//New Creates a new trigger instance for a given id
func (t *wsTriggerFactory) New(config *trigger.Config) trigger.Trigger {
	return &wsTrigger{metadata: t.metadata, config: config}
}

// udpTrigger is a stub for your Trigger implementation
type wsTrigger struct {
	metadata *trigger.Metadata
	runner   action.Runner
	config   *trigger.Config
	handlers []*trigger.Handler
}

// Metadata implements trigger.Trigger.Metadata
func (t *wsTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}



func (t *wsTrigger) Initialize(ctx trigger.InitContext) error {
	log.Debug("Initialize")
	t.handlers = ctx.GetHandlers()
	return nil

}

// Start implements trigger.Trigger.Start
func (t *wsTrigger) Start() error {


	// start the trigger
	log.Debug("Start")

	// Get parms
	wsUrl := t.config.GetSetting("url")

	// Parse the raw url string
	u, err := url.Parse(wsUrl)
	if err != nil {
		log.Errorf("Error parsing url %v", err)
		return nil
	}

  // Connect to Web Socket
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	dialer := websocket.Dialer{TLSClientConfig: tlsConfig, EnableCompression: true}
	conn, resp, err := dialer.Dial(u.String(), nil)
	if err != nil {
		log.Errorf("Handshake failed with status %d",resp.StatusCode)
		return nil
	}

	log.Debugf("WebSocket connection established to: [%s]",u.String())


	done := make(chan struct{})
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		defer conn.Close()
		defer close(done)
		for {
			log.Debug("event processing cycle starting")
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Errorf("read:", err)
				return
			}

			log.Debug("received message from websocket")

			var s = string(message)

			if strings.HasPrefix(s, "Response To") {
				log.Debug("acknowledgement received")
			} else {

				log.Info(s)

				log.Debug("event received")

				trgData := make(map[string]interface{})
				trgData["event"] = s
				for _, handler := range t.handlers {
					results, err := handler.Handle(context.Background(),trgData)
					if err != nil {
						log.Error("Error starting action: %v",err)
					}
					log.Debugf("Ran Handler: [%s]",handler)
					log.Debugf("Results [%v]",results)
				}


			}
			log.Debug("event processing cycle completed")
		}
	}()


	for {
		select {
		case t := <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Errorf("write:", err)
				return nil
			}
		case <-interrupt:
			log.Debug("interrupt")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Errorf("write close:", err)
				return nil
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			conn.Close()
			return nil
		}
	}


	return nil
}

// Stop implements trigger.Trigger.Start
func (t *wsTrigger) Stop() error {
	// stop the trigger
	return nil
}
