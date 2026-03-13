package bybit_connector

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type MessageHandler func(message string) error

func (b *WebSocket) handleIncomingMessages() {
	for {
		_, message, err := b.conn.ReadMessage()
		if err != nil {
			b.debug("Error reading: %v", err)
			b.isConnected = false
			return
		}

		if b.onMessage != nil {
			err := b.onMessage(string(message))
			if err != nil {
				b.debug("Error handling message: %v", err)
				return
			}
		}
	}
}

func (b *WebSocket) reconnect() bool {
	wssUrl := b.url
	if b.maxAliveTime != "" {
		wssUrl += "?max_alive_time=" + b.maxAliveTime
	}
	conn, _, err := websocket.DefaultDialer.Dial(wssUrl, nil)
	if err != nil {
		b.debug("Reconnect dial failed: %v", err)
		return false
	}
	b.conn = conn

	if b.requiresAuthentication() {
		if err = b.sendAuth(); err != nil {
			b.debug("Reconnect auth failed: %v", err)
			b.conn.Close()
			b.conn = nil
			return false
		}
	}
	return true
}

func (b *WebSocket) monitorConnection() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		select {
		case <-b.ctx.Done():
			return
		case <-ticker.C:
			if !b.isConnected {
				b.debug("Attempting to reconnect...")
				if b.reconnect() {
					b.isConnected = true
					go b.handleIncomingMessages()
				} else {
					b.debug("Reconnection failed")
				}
			}
		}
	}
}

func (b *WebSocket) SetMessageHandler(handler MessageHandler) {
	b.onMessage = handler
}

type WebSocket struct {
	conn         *websocket.Conn
	writeMu      sync.Mutex
	url          string
	apiKey       string
	apiSecret    string
	maxAliveTime string
	pingInterval int
	onMessage    MessageHandler
	ctx          context.Context
	cancel       context.CancelFunc
	isConnected  bool
	isDebug      bool
	Logger       *log.Logger
}

type WebsocketOption func(*WebSocket)

func WithPingInterval(pingInterval int) WebsocketOption {
	return func(c *WebSocket) {
		c.pingInterval = pingInterval
	}
}

func WithMaxAliveTime(maxAliveTime string) WebsocketOption {
	return func(c *WebSocket) {
		c.maxAliveTime = maxAliveTime
	}
}

// WithDebug print more details in log mode
func WithWsDebug(debug bool) WebsocketOption {
	return func(c *WebSocket) {
		c.isDebug = debug
	}
}

func NewBybitPrivateWebSocket(url, apiKey, apiSecret string, handler MessageHandler, options ...WebsocketOption) *WebSocket {
	c := &WebSocket{
		url:          url,
		apiKey:       apiKey,
		apiSecret:    apiSecret,
		maxAliveTime: "",
		pingInterval: 20,
		onMessage:    handler,
		Logger:       log.New(os.Stderr, Name, log.LstdFlags),
		isDebug:      false,
	}

	// Apply the provided options
	for _, opt := range options {
		opt(c)
	}

	return c
}

func NewBybitPublicWebSocket(url string, handler MessageHandler, options ...WebsocketOption) *WebSocket {
	c := &WebSocket{
		url:          url,
		maxAliveTime: "",
		pingInterval: 20,
		onMessage:    handler,
		Logger:       log.New(os.Stderr, Name, log.LstdFlags),
		isDebug:      false,
	}

	// Apply the provided options
	for _, opt := range options {
		opt(c)
	}

	return c
}

func (b *WebSocket) Connect() *WebSocket {
	var err error
	wssUrl := b.url
	if b.maxAliveTime != "" {
		wssUrl += "?max_alive_time=" + b.maxAliveTime
	}
	b.conn, _, err = websocket.DefaultDialer.Dial(wssUrl, nil)
	if err != nil {
		b.debug("Failed to connect: %v", err)
		return nil
	}

	if b.requiresAuthentication() {
		if err = b.sendAuth(); err != nil {
			b.debug("Failed Authentication: %v", err)
			b.conn.Close()
			b.conn = nil
			return nil
		}
	}
	b.isConnected = true

	b.ctx, b.cancel = context.WithCancel(context.Background())
	go b.handleIncomingMessages()
	go b.monitorConnection()
	go ping(b)

	return b
}

func (b *WebSocket) SendSubscription(args []string) (*WebSocket, error) {
	reqID := uuid.New().String()
	subMessage := map[string]interface{}{
		"req_id": reqID,
		"op":     "subscribe",
		"args":   args,
	}
	b.debug("subscribe msg: %v", subMessage["args"])
	if err := b.sendAsJson(subMessage); err != nil {
		b.debug("Failed to send subscription: %v", err)
		return b, err
	}
	b.debug("Subscription sent successfully.")
	return b, nil
}

// SendRequest sendRequest sends a custom request over the WebSocket connection.
func (b *WebSocket) SendRequest(op string, args map[string]interface{}, headers map[string]string, reqId ...string) (*WebSocket, error) {
	finalReqId := uuid.New().String()
	if len(reqId) > 0 && reqId[0] != "" {
		finalReqId = reqId[0]
	}

	request := map[string]interface{}{
		"reqId":  finalReqId,
		"header": headers,
		"op":     op,
		"args":   []interface{}{args},
	}
	b.debug("request headers: %v", request["header"])
	b.debug("request op channel: %v", request["op"])
	b.debug("request msg: %v", request["args"])
	if err := b.sendAsJson(request); err != nil {
		b.debug("Failed to send websocket trade request: %v", err)
		return b, err
	}
	b.debug("Successfully sent websocket trade request.")
	return b, nil
}

func (b *WebSocket) SendTradeRequest(tradeTruest map[string]interface{}) (*WebSocket, error) {
	b.debug("trade request headers: %v", tradeTruest["header"])
	b.debug("trade request op channel: %v", tradeTruest["op"])
	b.debug("trade request msg: %v", tradeTruest["args"])
	if err := b.sendAsJson(tradeTruest); err != nil {
		b.debug("Failed to send websocket trade request: %v", err)
		return b, err
	}
	b.debug("Successfully sent websocket trade request.")
	return b, nil
}

func ping(b *WebSocket) {
	if b.pingInterval <= 0 {
		b.debug("Ping interval is set to a non-positive value.")
		return
	}

	ticker := time.NewTicker(time.Duration(b.pingInterval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			currentTime := time.Now().Unix()
			pingMessage := map[string]string{
				"op":     "ping",
				"req_id": fmt.Sprintf("%d", currentTime),
			}
			jsonPingMessage, err := json.Marshal(pingMessage)
			if err != nil {
				b.debug("Failed to marshal ping message: %v", err)
				continue
			}
			b.writeMu.Lock()
			err = b.conn.WriteMessage(websocket.TextMessage, jsonPingMessage)
			b.writeMu.Unlock()
			if err != nil {
				b.debug("Failed to send ping: %v", err)
				return
			}
			b.debug("Ping sent with UTC time: %v", currentTime)

		case <-b.ctx.Done():
			b.debug("Ping context closed, stopping ping.")
			return
		}
	}
}

func (b *WebSocket) Disconnect() error {
	b.cancel()
	b.isConnected = false
	return b.conn.Close()
}

func (b *WebSocket) requiresAuthentication() bool {
	return b.url == WEBSOCKET_PRIVATE_MAINNET || b.url == WEBSOCKET_PRIVATE_TESTNET ||
		b.url == WEBSOCKET_TRADE_MAINNET || b.url == WEBSOCKET_TRADE_TESTNET ||
		b.url == WEBSOCKET_TRADE_DEMO || b.url == WEBSOCKET_PRIVATE_DEMO
	// v3 offline
	/*
		b.url == V3_CONTRACT_PRIVATE ||
			b.url == V3_UNIFIED_PRIVATE ||
			b.url == V3_SPOT_PRIVATE
	*/
}

func (b *WebSocket) sendAuth() error {
	// Get current Unix time in milliseconds
	expires := time.Now().UnixNano()/1e6 + 10000
	val := fmt.Sprintf("GET/realtime%d", expires)

	h := hmac.New(sha256.New, []byte(b.apiSecret))
	h.Write([]byte(val))

	// Convert to hexadecimal instead of base64
	signature := hex.EncodeToString(h.Sum(nil))
	b.debug("signature generated : " + signature)

	authMessage := map[string]interface{}{
		"req_id": uuid.New(),
		"op":     "auth",
		"args":   []interface{}{b.apiKey, expires, signature},
	}
	b.debug("auth args: %v", authMessage["args"])
	return b.sendAsJson(authMessage)
}

func (b *WebSocket) sendAsJson(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return b.send(string(data))
}

func (b *WebSocket) send(message string) error {
	b.writeMu.Lock()
	defer b.writeMu.Unlock()
	return b.conn.WriteMessage(websocket.TextMessage, []byte(message))
}

func (c *WebSocket) debug(format string, v ...interface{}) {
	if c.isDebug {
		c.Logger.Printf(format, v...)
	}
}
