package mqttconnection

import (
	b64 "encoding/base64"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/connection"
	"github.com/project-flogo/core/support/log"
	"strings"
	"time"
)

type Settings struct {
	Name          string `md:"name,required"`
	Description   string `md:"description"`
	ConnectionUri string `md:"connectionURI"`
	ClientID      string `md:"clientID"`
	KeepAlive     int    `md:"keepAlive"`

	//	DocsMetadata  string `md:"DocsMetadata"`
	CredType   string `md:"credType"`
	UserName   string `md:"username"`
	Password   string `md:"password"`
	Ssl        bool   `md:"ssl"`
	TrustCert  string `md:"trustCert"`
	ClientKey  string `md:"clientKey"`
	ClientCert string `md:"clientCert"`
	KeyPass    string `md:"keyPassword"`
	X509       bool   `md:"x509"`
}

type mqttFactory struct {
}

// MqttConfigManager Structure
type MqttSharedConfigManager struct {
	config *Settings
	name   string
	client *mqtt.Client
}

var logger = log.ChildLogger(log.RootLogger(), "mqtt-connection")
var factory = &mqttFactory{}

func init() {
	err := connection.RegisterManagerFactory(factory)
	if err != nil {
		panic(err)
	}
}

func (*mqttFactory) Type() string {
	return "mqtt"
}

func (*mqttFactory) NewManager(settings map[string]interface{}) (connection.Manager, error) {

	sharedConn := &MqttSharedConfigManager{}
	var err error
	sharedConn.config, err = getmqttClientConfig(settings)
	if err != nil {
		return nil, err
	}
	if sharedConn.client != nil {
		return sharedConn, nil
	}
	opts := mqtt.NewClientOptions()
	ssl := sharedConn.config.Ssl
	credType := sharedConn.config.CredType

	opts.AddBroker(sharedConn.config.ConnectionUri)
	opts.SetClientID(sharedConn.config.ClientID)
	opts.SetKeepAlive(time.Duration(sharedConn.config.KeepAlive))

	if credType != "None" {
		userName := sharedConn.config.UserName
		password := sharedConn.config.Password
		opts.SetUsername(userName)
		opts.SetPassword(password)
	}

	if ssl {

	}

	client := mqtt.NewClient(opts)
	if err != nil {
		return nil, err
	}

	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sharedConn.client = &client
	return sharedConn, nil
}

// Type of SharedConfigManager
func (k *MqttSharedConfigManager) Type() string {
	return "mqtt"
}

// GetConnection
func (k *MqttSharedConfigManager) GetConnection() interface{} {
	return k.client
}

// ReleaseConnection
func (k *MqttSharedConfigManager) ReleaseConnection(connection interface{}) {

}

// Start connection manager
func (k *MqttSharedConfigManager) Start() error {
	return nil
}

// Stop connection manager
func (k *MqttSharedConfigManager) Stop() error {
	return nil
}

// GetSharedConfiguration function to return Mqtt connection manager
func GetSharedConfiguration(conn interface{}) (connection.Manager, error) {
	var cManager connection.Manager
	var err error
	cManager, err = coerce.ToConnection(conn)
	if err != nil {
		return nil, err
	}
	return cManager, nil
}

func getmqttClientConfig(settings map[string]interface{}) (*Settings, error) {
	connectionConfig := &Settings{}
	err := metadata.MapToStruct(settings, connectionConfig, false)
	if err != nil {
		return nil, err
	}
	return connectionConfig, nil
}

// parse cert

func parseCert(cert string) string {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(cert), &m)
	if err != nil {
		logger.Errorf("=======Error Parsing Certificate for SSL handshake=====", err)
	}
	content := m["content"].(string)
	lastBin := strings.LastIndex(content, "base64,")
	sEnc := content[lastBin+7 : len(content)]
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	return (string(sDec))
}
