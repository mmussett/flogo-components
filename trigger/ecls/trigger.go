package ecls

import (
	"context"
	"encoding/json"
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	"crypto/tls"
	"github.com/gorilla/websocket"
	"net/url"

	"os"
	"os/signal"
	"strings"
	"time"
)

type ECLSMessage struct {
	Data []struct {
		RequestHostName               string    `json:"request_host_name"`
		SrcIP                         string    `json:"src_ip"`
		RequestUUID                   string    `json:"request_uuid"`
		HTTPMethod                    string    `json:"http_method"`
		URI                           string    `json:"uri"`
		HTTPVersion                   string    `json:"http_version"`
		Bytes                         string    `json:"bytes"`
		HTTPStatusCode                string    `json:"http_status_code"`
		Referrer                      string    `json:"referrer"`
		UserAgent                     string    `json:"user_agent"`
		RequestID                     string    `json:"request_id"`
		RequestTime                   time.Time `json:"request_time"`
		APIKey                        string    `json:"api_key"`
		ServiceID                     string    `json:"service_id"`
		TrafficManager                string    `json:"traffic_manager"`
		APIMethodName                 string    `json:"api_method_name"`
		CacheHit                      string    `json:"cache_hit"`
		TrafficManagerErrorCode       string    `json:"traffic_manager_error_code"`
		TotalRequestExecTime          string    `json:"total_request_exec_time"`
		RemoteTotalTime               string    `json:"remote_total_time"`
		ConnectTime                   string    `json:"connect_time"`
		PreTransferTime               string    `json:"pre_transfer_time"`
		OauthAccessToken              string    `json:"oauth_access_token"`
		SslEnabled                    string    `json:"ssl_enabled"`
		QuotaValue                    string    `json:"quota_value"`
		QPSThrottleValue              string    `json:"qps_throttle_value"`
		ClientTransferTime            string    `json:"client_transfer_time"`
		ServiceName                   string    `json:"service_name"`
		ResponseString                string    `json:"response_string"`
		PlanName                      string    `json:"plan_name"`
		PlanUUID                      string    `json:"plan_uuid"`
		EndpointName                  string    `json:"endpoint_name"`
		PackageName                   string    `json:"package_name"`
		PackageUUID                   string    `json:"package_uuid"`
		ServiceDefinitionEndpointUUID string    `json:"service_definition_endpoint_uuid"`
		LogType                       string    `json:"log_type"`
		IngestionTime                 string    `json:"ingestion_time"`
	} `json:"data"`
}

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
		log.Errorf("Handshake failed with status %d", resp.StatusCode)
		return nil
	}

	log.Debugf("WebSocket connection established to: [%s]", u.String())

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

				var ecls_message ECLSMessage

				err := json.Unmarshal([]byte(s), &ecls_message)
				if err == nil {


					log.Debug("event received")

					trgData := make(map[string]interface{})
					trgData["api_key"] = ecls_message.Data[0].APIKey
					trgData["api_method_name"] = ecls_message.Data[0].APIMethodName
					trgData["bytes"] = ecls_message.Data[0].Bytes
					trgData["cache_hit"] = ecls_message.Data[0].CacheHit
					trgData["client_transfer_time"] = ecls_message.Data[0].ClientTransferTime
					trgData["connect_time"] = ecls_message.Data[0].ConnectTime
					trgData["endpoint_name"] = ecls_message.Data[0].EndpointName
					trgData["http_method"] = ecls_message.Data[0].HTTPMethod
					trgData["http_status_code"] = ecls_message.Data[0].HTTPStatusCode
					trgData["http_version"] = ecls_message.Data[0].HTTPVersion
					trgData["oauth_access_token"] = ecls_message.Data[0].OauthAccessToken
					trgData["package_name"] = ecls_message.Data[0].PackageName
					trgData["package_uuid"] = ecls_message.Data[0].PackageUUID
					trgData["plan_name"] = ecls_message.Data[0].PlanName
					trgData["plan_uuid"] = ecls_message.Data[0].PlanUUID
					trgData["pre_transfer_time"] = ecls_message.Data[0].PreTransferTime
					trgData["qps_throttle_value"] = ecls_message.Data[0].QPSThrottleValue
					trgData["quota_value"] = ecls_message.Data[0].QuotaValue
					trgData["referrer"] = ecls_message.Data[0].Referrer
					trgData["remote_total_time"] = ecls_message.Data[0].RemoteTotalTime
					trgData["request_host_name"] = ecls_message.Data[0].RequestHostName
					trgData["request_id"] = ecls_message.Data[0].RequestID
					trgData["request_time"] = ecls_message.Data[0].RequestTime
					trgData["request_uuid"] = ecls_message.Data[0].RequestUUID
					trgData["response_string"] = ecls_message.Data[0].ResponseString
					trgData["service_definition_endpoint_uuid"] = ecls_message.Data[0].ServiceDefinitionEndpointUUID
					trgData["service_id"] = ecls_message.Data[0].ServiceID
					trgData["service_name"] = ecls_message.Data[0].ServiceName
					trgData["src_ip"] = ecls_message.Data[0].SrcIP
					trgData["ssl_enabled"] = ecls_message.Data[0].SslEnabled
					trgData["total_request_exec_time"] = ecls_message.Data[0].TotalRequestExecTime
					trgData["traffic_manager"] = ecls_message.Data[0].TrafficManager
					trgData["traffic_manager_error_code"] = ecls_message.Data[0].TrafficManagerErrorCode
					trgData["uri"] = ecls_message.Data[0].URI
					trgData["user_agent"] = ecls_message.Data[0].UserAgent

					for _, handler := range t.handlers {
						results, err := handler.Handle(context.Background(), trgData)
						if err != nil {
							log.Error("Error starting action: %v", err)
						}
						log.Debugf("Ran Handler: [%s]", handler)
						log.Debugf("Results [%v]", results)
					}
				} else {
					log.Error("Unable to unmarshal ECLS event: %v", err)
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
