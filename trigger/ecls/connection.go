package ecls

import (
	"crypto/tls"
	"fmt"

	"net/url"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/prometheus/common/log"
)

type WebSocketConnection struct {
	conn *websocket.Conn
}

func (c *WebSocketConnection) Connection() *websocket.Conn {
	return c.conn
}

func (c *WebSocketConnection) Stop() error {
	return c.conn.Close()
}

func NewWebSocketConnection(wsUrl string) (*WebSocketConnection, error) {

	log.Debugf("WebSocket URL: [%v]", wsUrl)

	newConn := &WebSocketConnection{}

	err := validateUrl(wsUrl)
	if err != nil {
		return nil, fmt.Errorf("url [%s] format invalid for reason: [%v]", wsUrl, err)
	}

	// Parse the raw url string
	u, err := url.Parse(wsUrl)
	if err != nil {
		return nil, fmt.Errorf("error parsing url %v", err)
	}

	// Connect to Web Socket
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	dialer := websocket.Dialer{TLSClientConfig: tlsConfig, EnableCompression: true}

	c, resp, err := dialer.Dial(u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("handshake failed with status %d", resp.StatusCode)
	}

	newConn.conn = c
	return newConn, nil
}

// validateUrl to make sure it starts with wss:// or ws://
func validateUrl(url string) error {

	if strings.HasPrefix(url, "wss://") || strings.HasPrefix(url, "ws://") {
		return nil
	} else {
		return fmt.Errorf("not a websoocket url")
	}
}
