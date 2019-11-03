package ecls

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/trigger"
	"github.com/prometheus/common/log"

	"reflect"

	"github.com/gorilla/websocket"

	"os"
	"os/signal"
	"strings"
	"time"
)

type Message struct {
	Data []struct {
		RequestHostName               string `json:"request_host_name"`
		SrcIP                         string `json:"src_ip"`
		RequestUUID                   string `json:"request_uuid"`
		HTTPMethod                    string `json:"http_method"`
		URI                           string `json:"uri"`
		HTTPVersion                   string `json:"http_version"`
		Bytes                         string `json:"bytes"`
		HTTPStatusCode                string `json:"http_status_code"`
		Referrer                      string `json:"referrer"`
		UserAgent                     string `json:"user_agent"`
		RequestID                     string `json:"request_id"`
		RequestTime                   string `json:"request_time"`
		APIKey                        string `json:"api_key"`
		ServiceID                     string `json:"service_id"`
		TrafficManager                string `json:"traffic_manager"`
		APIMethodName                 string `json:"api_method_name"`
		CacheHit                      string `json:"cache_hit"`
		TrafficManagerErrorCode       string `json:"traffic_manager_error_code"`
		TotalRequestExecTime          string `json:"total_request_exec_time"`
		RemoteTotalTime               string `json:"remote_total_time"`
		ConnectTime                   string `json:"connect_time"`
		PreTransferTime               string `json:"pre_transfer_time"`
		OauthAccessToken              string `json:"oauth_access_token"`
		SslEnabled                    string `json:"ssl_enabled"`
		QuotaValue                    string `json:"quota_value"`
		QPSThrottleValue              string `json:"qps_throttle_value"`
		ClientTransferTime            string `json:"client_transfer_time"`
		ServiceName                   string `json:"service_name"`
		ResponseString                string `json:"response_string"`
		PlanName                      string `json:"plan_name"`
		PlanUUID                      string `json:"plan_uuid"`
		EndpointName                  string `json:"endpoint_name"`
		PackageName                   string `json:"package_name"`
		PackageUUID                   string `json:"package_uuid"`
		ServiceDefinitionEndpointUUID string `json:"service_definition_endpoint_uuid"`
		LogType                       string `json:"log_type"`
		IngestionTime                 string `json:"ingestion_time"`
	} `json:"data"`
}

var triggerMd = trigger.NewMetadata(&Settings{}, &Output{})

func init() {
	_ = trigger.Register(&Trigger{}, &Factory{})
}

type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// New implements trigger.Factory.New
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)

	if err != nil {
		return nil, err
	}

	return &Trigger{id: config.Id, settings: s}, nil
}

type Trigger struct {
	settings   *Settings
	id         string
	Connection *WebSocketConnection
	Handlers   []trigger.Handler
}

func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	c, err := NewWebSocketConnection(t.settings.Url)

	if err != nil {
		ctx.Logger().Error(err)
		return err
	}

	t.Connection = c

	for _, handler := range ctx.GetHandlers() {
		t.Handlers = append(t.Handlers, handler)
	}

	return nil
}

func (t *Trigger) Start() error {

	return t.startHandlers()
}

// Stop implements util.Managed.Stop
func (t *Trigger) Stop() error {
	_ = t.Connection.Stop()

	return nil
}

func (t *Trigger) startHandlers() error {

	// start the trigger
	log.Debug("Start")

	/*	// Get parms
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
	*/

	conn := t.Connection.conn

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
				log.Error("read:", err)
				return
			}

			log.Debug("received message from websocket")

			var s = string(message)

			if strings.HasPrefix(s, "Response To") {
				log.Debug("acknowledgement received")
			} else {

				log.Info(s)

				var eclsMessage Message

				err := json.Unmarshal([]byte(s), &eclsMessage)

				eclsMessage.flatten()

				if err == nil {

					log.Debug("event received")

					trgData := make(map[string]interface{})
					trgData["api_key"] = eclsMessage.Data[0].APIKey
					trgData["api_method_name"] = eclsMessage.Data[0].APIMethodName
					trgData["bytes"] = eclsMessage.Data[0].Bytes
					trgData["cache_hit"] = eclsMessage.Data[0].CacheHit
					trgData["client_transfer_time"] = eclsMessage.Data[0].ClientTransferTime
					trgData["connect_time"] = eclsMessage.Data[0].ConnectTime
					trgData["endpoint_name"] = eclsMessage.Data[0].EndpointName
					trgData["http_method"] = eclsMessage.Data[0].HTTPMethod
					trgData["http_status_code"] = eclsMessage.Data[0].HTTPStatusCode
					trgData["http_version"] = eclsMessage.Data[0].HTTPVersion
					trgData["oauth_access_token"] = eclsMessage.Data[0].OauthAccessToken
					trgData["package_name"] = eclsMessage.Data[0].PackageName
					trgData["package_uuid"] = eclsMessage.Data[0].PackageUUID
					trgData["plan_name"] = eclsMessage.Data[0].PlanName
					trgData["plan_uuid"] = eclsMessage.Data[0].PlanUUID
					trgData["pre_transfer_time"] = eclsMessage.Data[0].PreTransferTime
					trgData["qps_throttle_value"] = eclsMessage.Data[0].QPSThrottleValue
					trgData["quota_value"] = eclsMessage.Data[0].QuotaValue
					trgData["referrer"] = eclsMessage.Data[0].Referrer
					trgData["remote_total_time"] = eclsMessage.Data[0].RemoteTotalTime
					trgData["request_host_name"] = eclsMessage.Data[0].RequestHostName
					trgData["request_id"] = eclsMessage.Data[0].RequestID
					trgData["request_time"] = eclsMessage.Data[0].RequestTime
					trgData["request_uuid"] = eclsMessage.Data[0].RequestUUID
					trgData["response_string"] = eclsMessage.Data[0].ResponseString
					trgData["service_definition_endpoint_uuid"] = eclsMessage.Data[0].ServiceDefinitionEndpointUUID
					trgData["service_id"] = eclsMessage.Data[0].ServiceID
					trgData["service_name"] = eclsMessage.Data[0].ServiceName
					trgData["src_ip"] = eclsMessage.Data[0].SrcIP
					trgData["ssl_enabled"] = eclsMessage.Data[0].SslEnabled
					trgData["total_request_exec_time"] = eclsMessage.Data[0].TotalRequestExecTime
					trgData["traffic_manager"] = eclsMessage.Data[0].TrafficManager
					trgData["traffic_manager_error_code"] = eclsMessage.Data[0].TrafficManagerErrorCode
					trgData["uri"] = eclsMessage.Data[0].URI
					trgData["user_agent"] = eclsMessage.Data[0].UserAgent
					trgData["log_type"] = eclsMessage.Data[0].LogType
					trgData["ingestion_time"] = eclsMessage.Data[0].IngestionTime
					trgData["asCSV"] = eclsMessage.flatten()
					trgData["asObject"] = s

					for _, handler := range t.Handlers {
						results, err := handler.Handle(context.Background(), trgData)
						if err != nil {
							log.Errorf("Error starting action: %v", err)
						}
						log.Debugf("Ran Handler: [%s]", handler)
						log.Debugf("Results [%v]", results)
					}
				} else {
					log.Errorf("Unable to unmarshal ECLS event: %v", err)
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
				log.Error("write:", err)
				return nil
			}
		case <-interrupt:
			log.Debug("interrupt")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Error("write close:", err)
				return nil
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			_ = conn.Close()
			return nil
		}
	}

}

func (t *Message) flatten() string {

	v := reflect.ValueOf(t.Data[0])

	fields := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		v := fmt.Sprintf("\"%s\"", f.Interface())
		fields[i] = v

	}

	r := strings.Join(fields, ",")

	return r

}
