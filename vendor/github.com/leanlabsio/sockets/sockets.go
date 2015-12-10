// Copyright 2014 Beat Richartz
// Copyright 2014 The Macaron Authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package sockets is a middleware that provides WebSockets channels binding for Macaron.
package sockets

// NOTE: last sync 97a57b4 on Jul 7, 2014.

import (
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/gorilla/websocket"
	"gopkg.in/macaron.v1"
)

type LogLevel int

const (
	// Log levels 0-4. Use to set the log level you wish to go for.
	LEVEL_ERROR LogLevel = iota
	LEVEL_WARNING
	LEVEL_INFO
	LEVEL_DEBUG

	// Sensible defaults for the socket
	defaultLogLevel                = LEVEL_INFO
	defaultWriteWait               = 60 * time.Second
	defaultPongWait                = 60 * time.Second
	defaultPingPeriod              = (defaultPongWait * 8 / 10)
	defaultMaxMessageSize    int64 = 65536
	defaultSendChannelBuffer       = 10
	defaultRecvChannelBuffer       = 10
)

type Options struct {
	// The logger to use for socket logging
	Logger *log.Logger

	// The LogLevel for socket logging, goes from 0 (Error) to 3 (Debug)
	LogLevel LogLevel

	// Set to true if you want to skip logging
	SkipLogging bool

	// The time to wait between writes before timing out the connection
	// When this is a zero value time instance, write will never time out
	WriteWait time.Duration

	// The time to wait at maximum between receiving pings from the client.
	PongWait time.Duration

	// The time to wait between sending pings to the client
	PingPeriod time.Duration

	// The maximum messages size for receiving and sending in bytes
	MaxMessageSize int64

	// The send channel buffer
	SendChannelBuffer int

	// The receiving channel buffer
	RecvChannelBuffer int
}

type Connection struct {
	*Options

	// The websocket connection
	ws *websocket.Conn

	// The remote Address of the client using this connection. Cached on the
	// connection for logging.
	remoteAddr net.Addr

	// The error channel is given the error object as soon as an error occurs
	// either sending or receiving values from the websocket. This channel gets
	// mapped for the next handler to use.
	Error chan error

	// The disconnect channel is for listening for disconnects from the next handler.
	// Any sends to the disconnect channel lead to disconnecting the socket with the
	// given closing message. This channel gets mapped for the next
	// handler to use.
	Disconnect chan int

	// The done channel gets called only when the connection
	// has been successfully disconnected. Any sends to the disconnect
	// channel are currently ignored. This channel gets mapped for the next
	// handler to use.
	Done chan bool

	// The internal disconnect channel. Sending on this channel will lead to the handlers and
	// the connection closing.
	disconnect chan error

	// The disconnect send channel. Sending on this channel will lead to the send handler and
	// closing.
	disconnectSend chan bool

	// the ticker for pinging the client.
	ticker *time.Ticker
}

type Binding interface {
	Close(int) error
	recv()
	send()
	setSocketOptions()
	mapChannels(*macaron.Context)
	mapDefaultChannels(*macaron.Context)
	disconnectChannel() chan error
	DisconnectChannel() chan int
	ErrorChannel() chan error
}

// Message Connection connects a websocket message connection to a string
// channel.
type MessageConnection struct {
	*Connection

	// Sender is the string channel used for sending out strings to the client.
	// This channel gets mapped for the next handler to use and is asynchronous
	// unless the SendChannelBuffer is set to 0.
	Sender chan string

	// Receiver is the string channel used for receiving strings from the client.
	// This channel gets mapped for the next handler to use and is asynchronous
	// unless the RecvChannelBuffer is set to 0.
	Receiver chan string
}

// Message Connection connects a websocket message connection to a reflect.Value
// channel.
type JSONConnection struct {
	*Connection

	// Sender is the channel used for sending out JSON to the client.
	// This channel gets mapped for the next handler to use with the right type
	// and is asynchronous unless the SendChannelBuffer is set to 0.
	Sender reflect.Value

	// Receiver is the string channel used for receiving JSON from the client.
	// This channel gets mapped for the next handler to use with the right type
	// and is asynchronous unless the RecvChannelBuffer is set to 0.
	Receiver reflect.Value
}

// Messages returns a websocket handling middleware. It can only be used
// in handlers for HTTP GET.
// IMPORTANT: The last handler in your handler chain must block in order for the
// connection to be kept alive.
// It maps four channels for you to use in the follow-up Handler(s):
// - A receiving string channel (<-chan string) on which you will
//   receive all incoming strings from the client
// - A sending string channel (chan<- string) on which you will be
//   able to send strings to the client.
// - A receiving error channel (<-chan error) on which you will receive
//   errors occurring while sending & receiving
// - A receiving disconnect channel  (<-chan bool) on which you will receive
//   a message only if the connection is about to be closed following an
//   error or a client disconnect.
// - A sending done channel  (chan<- bool) on which you can send as soon as you wish
//   to disconnect the connection.
// The middleware handles the following for you:
// - Checking the request for cross origin access
// - Doing the websocket handshake
// - Setting sensible options for the Gorilla websocket connection
// - Starting and terminating the necessary goroutines
// An optional sockets.Options object can be passed to Messages to overwrite
// default options mentioned in the documentation of the Options object.
func Messages(options ...*Options) macaron.Handler {
	return makeHandler("", newOptions(options))
}

// JSON returns a websocket handling middleware. It can only be used
// in handlers for HTTP GET.
// IMPORTANT: The last handler in your handler chain must block in order for the
// connection to be kept alive.
// It accepts an empty struct it will copy and try to populate
// with data received from the client using the JSON Marshaler, as well
// as it will serialize your structs to JSON and send them to the client.
// For the following, it is assumed you passed a struct named Message
// to the handler.
// It maps four channels for you to use in the follow-up Handler(s):
// - A receiving string channel (<-chan *Message) on which you will
//   receive all incoming structs from the client
// - A sending string channel (chan<- *Message) on which you will be
//   able to send structs to the client.
// - A receiving error channel (<-chan error) on which you will receive
//   errors occurring while sending & receiving
// - A receiving disconnect channel  (<-chan bool) on which you will receive
//   a message only if the connection is about to be closed following an
//   error or a client disconnect.
// - A sending done channel  (chan<- bool) on which you can send as soon as you wish
//   to disconnect the connection.
// The middleware handles the following for you:
// - Checking the request for cross origin access
// - Doing the websocket handshake
// - Setting sensible options for the Gorilla websocket connection
// - Starting and terminating the necessary goroutines
// An optional sockets.Options object can be passed to Messages to overwrite
// default options mentioned in the documentation of the Options object.
func JSON(bindStruct interface{}, options ...*Options) macaron.Handler {
	return makeHandler(bindStruct, newOptions(options))
}

// Generates a handler from an interface
func makeHandler(binding interface{}, o *Options) macaron.Handler {

	return func(ctx *macaron.Context) {
		// Upgrade the request to a websocket connection
		ws, status, err := upgradeRequest(ctx.Resp, ctx.Req.Request, o)
		if err != nil {
			ctx.Resp.WriteHeader(status)
			ctx.Resp.Write([]byte(err.Error()))
			return
		}

		// Set up the connection
		c := newBinding(binding, ws, o)

		// Set the options for the gorilla websocket package
		c.setSocketOptions()

		// Map the sending and receiving channels
		c.mapChannels(ctx)

		// Map the Channels <-chan error, <-chan bool and chan<- bool
		c.mapDefaultChannels(ctx)

		// start the send and receive goroutines
		go c.send()
		go c.recv()
		go waitForDisconnect(c)

		// call the next handler, which must block
		ctx.Next()
	}
}

// Log Level to strings slice
var LogLevelStrings = []string{"Error", "Warning", "Info", "Debug"}

// The options logger is only directly used while setting up the connection
// With the default logger, it logs in the format [socket][client remote address] log message
func (o *Options) log(message string, level LogLevel, logVars ...interface{}) {
	if level <= o.LogLevel && !o.SkipLogging {
		o.Logger.Printf("[%s] [%s] "+message, append([]interface{}{LogLevelStrings[level]}, logVars...)...)
	}
}

// The connection logger writes to the option logger using the cached remote address
// for this connection
func (c *Connection) log(message string, level LogLevel, logVars ...interface{}) {
	if level <= c.LogLevel {
		c.Options.log(message, level, append([]interface{}{c.remoteAddr}, logVars...)...)
	}
}

// Set the gorilla websocket handler options according to given options and set a default pong
// handler to keep the connection alive
func (c *Connection) setSocketOptions() {
	c.ws.SetReadLimit(c.MaxMessageSize)
	c.keepAlive()
	c.ws.SetPongHandler(func(string) error {
		c.log("Received Pong from Client", LEVEL_DEBUG)
		c.keepAlive()
		return nil
	})
}

// Helper method to map default channels in the context
// Map the Error Channel to a <-chan error for the next Handler(s)
// Map the Disconnect Channel to a chan<- bool for the next Handler(s)
// Map the Done Channel to a <-chan bool for the next Handler(s)
func (c *Connection) mapDefaultChannels(ctx *macaron.Context) {
	ctx.Set(reflect.ChanOf(reflect.RecvDir, reflect.TypeOf(c.Error).Elem()), reflect.ValueOf(c.Error))
	ctx.Set(reflect.ChanOf(reflect.SendDir, reflect.TypeOf(c.Disconnect).Elem()), reflect.ValueOf(c.Disconnect))
	ctx.Set(reflect.ChanOf(reflect.RecvDir, reflect.TypeOf(c.Done).Elem()), reflect.ValueOf(c.Done))
}

// Close the Base connection. Closes the send Handler and all channels used
// Since all channels are either internal or channels this middleware is sending on.
func (c *Connection) Close(closeCode int) error {
	c.disconnectSend <- true
	//TODO look for a better way to unblock the reader
	c.ws.SetReadDeadline(time.Now())

	// Send close message to the client
	c.log("Sending close message to client", LEVEL_DEBUG)
	c.ws.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(closeCode, ""), time.Now().Add(c.WriteWait))

	// If the connection can not be closed, return the error
	c.log("Closing websocket connection", LEVEL_DEBUG)
	if err := c.ws.Close(); err != nil {
		c.log("Connection could not be closed: %s", LEVEL_ERROR, err.Error())
		return err
	}

	// Send disconnect message to the next handler
	c.log("Sending disconnect to handler", LEVEL_DEBUG)
	c.Done <- true

	// Close disconnect and error channels this connection was sending on
	close(c.Done)
	close(c.Error)

	return nil
}

// Ping the client through the websocket
func (c *Connection) ping() error {
	c.log("Pinging socket", LEVEL_DEBUG)
	return c.ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(c.WriteWait))
}

// Start the ticker used for pinging the client
func (c *Connection) startTicker() {
	c.log("Pinging every %v, first at %v", LEVEL_DEBUG, c.PingPeriod, time.Now().Add(c.PingPeriod))
	c.ticker = time.NewTicker(c.PingPeriod)
}

// Stop the ticker used for pinging the client
func (c *Connection) stopTicker() {
	c.log("Stopped pinging socket", LEVEL_DEBUG)
	c.ticker.Stop()
}

// Keep the connection alive by refreshing the deadlines.
func (c *Connection) keepAlive() {
	c.log("Setting read deadline to %v", LEVEL_DEBUG, time.Now().Add(c.PongWait))
	c.ws.SetReadDeadline(time.Now().Add(c.PongWait))
	if c.WriteWait == 0 {
		c.log("Write deadline set to 0, will never expire", LEVEL_DEBUG)
		c.ws.SetWriteDeadline(time.Time{})
	} else {
		c.log("Setting write deadline to %v", LEVEL_DEBUG, time.Now().Add(c.WriteWait))
		c.ws.SetWriteDeadline(time.Now().Add(c.WriteWait))
	}
}

func (c *Connection) disconnectChannel() chan error {
	return c.disconnect
}

func (c *Connection) DisconnectChannel() chan int {
	return c.Disconnect
}

func (c *Connection) ErrorChannel() chan error {
	return c.Error
}

// Close the Message connection. Closes the send goroutine and all channels used
// Except for the send channel, since it should be closed by the handler sending on it.
func (c *MessageConnection) Close(closeCode int) error {
	// Call close on the base connection
	c.log("Closing websocket connection", LEVEL_DEBUG)
	err := c.Connection.Close(closeCode)

	if err != nil {
		return err
	}

	// Do not close the receiver here since it would send nil
	// Just let go
	c.log("Connection closed", LEVEL_INFO)

	return nil
}

// Write the message to the websocket, also keeping the connection alive
func (c *MessageConnection) write(mt int, payload string) error {
	c.keepAlive()
	return c.ws.WriteMessage(mt, []byte(payload))
}

// Send handler for the message connection. Starts a goroutine
// Listening on the sender channel and writing received strings
// to the websocket.
func (c *MessageConnection) send() {
	// Start the ticker and defer stopping it and decrementing the
	// wait group counter.
	c.startTicker()
	defer func() {
		c.stopTicker()
		c.log("Goroutine sending to websocket has been closed", LEVEL_DEBUG)
	}()

	for {
		select {
		// Receiving a message from the next handler
		case message, ok := <-c.Sender:
			if !ok {
				c.log("Sender channel has been closed", LEVEL_ERROR)
				c.disconnect <- errors.New("Sender channel has been closed")
				return
			}
			// Write the message as a byte array to the socket
			c.log("Writing %s to socket", LEVEL_DEBUG, message)
			if err := c.write(websocket.TextMessage, message); err != nil {
				c.log("Error writing to socket: %s", LEVEL_ERROR, err)
				c.disconnect <- err
				return
			}

			c.keepAlive()
		// Ping the client
		case <-c.ticker.C:
			err := c.ping()
			c.log("%s", LEVEL_DEBUG, err)
			if err := c.ping(); err != nil {
				c.log("Error pinging socket: %s", LEVEL_ERROR, err)
				c.disconnect <- err
				return
			}

		// Receiving disconnectSend from the closing Connection
		case <-c.disconnectSend:
			return
		}
	}
}

func (c *MessageConnection) recv() {
	// Defer decrementing the wait group counter and closing the connection
	defer func() {
		c.log("Goroutine receiving from websocket has been closed", LEVEL_DEBUG)
	}()

	for {
		// Read a message from the client
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			c.log("Error reading from socket: %s", LEVEL_ERROR, err)
			c.disconnect <- err
			return
		}
		// Send the message as a string to the next handler
		c.log("Read message from socket, %s", LEVEL_DEBUG, string(message))
		c.Receiver <- string(message)
		c.keepAlive()
	}
}

// Map the Receiver to a chan<- string for the next Handler(s)
// Map the Receiver to a <-chan string for the next Handler(s)
func (c *MessageConnection) mapChannels(ctx *macaron.Context) {
	ctx.Set(reflect.ChanOf(reflect.SendDir, reflect.TypeOf(c.Sender).Elem()), reflect.ValueOf(c.Sender))
	ctx.Set(reflect.ChanOf(reflect.RecvDir, reflect.TypeOf(c.Receiver).Elem()), reflect.ValueOf(c.Receiver))
}

// Close the JSON connection. Closes the send goroutine and all channels used
// Except for the send channel, since it should be closed by the handler sending on it.
func (c *JSONConnection) Close(closeCode int) error {
	// Call close on the base connection
	c.log("Closing websocket connection", LEVEL_DEBUG)
	err := c.Connection.Close(closeCode)
	if err != nil {
		return err
	}

	// Do not close the receiver here since it would send nil
	// Just let go
	c.log("Connection closed", LEVEL_INFO)

	return nil
}

var (
	senderSend     = 0
	tickerTick     = 1
	disconnectSend = 2
)

func (c *JSONConnection) send() {
	// Start the ticker and defer stopping it and decrementing the
	// wait group counter.
	c.startTicker()
	defer func() {
		c.stopTicker()
		c.log("Goroutine sending to websocket has been closed", LEVEL_DEBUG)
	}()

	// Creating the select cases for the channel select
	cases := make([]reflect.SelectCase, 3)

	// Case 0 listens on the sender, equals: case <-c.Sender:
	cases[senderSend] = reflect.SelectCase{reflect.SelectRecv, c.Sender, reflect.ValueOf(nil)}

	// Case 1 listens on the timer channel, equals: case <-c.ticker.C:
	cases[tickerTick] = reflect.SelectCase{reflect.SelectRecv, reflect.ValueOf(c.ticker.C), reflect.ValueOf(nil)}

	// Case 2 listens on the disconnectSend channel, equals: case <-disconnectSend:
	cases[disconnectSend] = reflect.SelectCase{reflect.SelectRecv, reflect.ValueOf(c.disconnectSend), reflect.ValueOf(nil)}

	for {
		chosen, message, ok := reflect.Select(cases)
		switch chosen {
		// Receiving a message from the next handler
		case senderSend:
			if !ok {
				c.log("Sender channel has been closed", LEVEL_ERROR)
				c.disconnect <- errors.New("Sender channel has been closed")
				return
			}
			c.log("Writing %v: %v to socket", LEVEL_DEBUG, message.Type(), message.Interface())
			if err := c.ws.WriteJSON(message.Interface()); err != nil {
				c.log("Error writing to socket: %s", LEVEL_ERROR, err)
				c.disconnect <- err
				break
			}
			c.keepAlive()
		// Pinging the client
		case tickerTick:
			if err := c.ping(); err != nil {
				c.log("Error pinging socket: %s", LEVEL_ERROR, err)
				c.disconnect <- err
				return
			}
		// Received disconnectSend from the closing connection
		case disconnectSend:
			return
		}
	}
}

func (c *JSONConnection) recv() {
	for {
		message := c.newOfType()

		err := c.ws.ReadJSON(message.Interface())
		if err != nil {
			c.log("Error reading from socket: %s", LEVEL_ERROR, err)
			c.disconnect <- err
			break
		}

		// Send the message to the next handler
		c.log("Read message from socket: %v: %v", LEVEL_DEBUG, message.Type(), message.Interface())
		c.Receiver.Send(message)
	}

	c.log("Goroutine receiving from websocket has been closed", LEVEL_DEBUG)
}

// Creates a new empty message of the given struct type
func (c *JSONConnection) newOfType() reflect.Value {
	return reflect.New(c.Sender.Type().Elem().Elem())
}

// Map the Sender to a chan<- *Message for the next Handler(s)
// Map the Receiver to a <-chan *Message for the next Handler(s)
func (c *JSONConnection) mapChannels(ctx *macaron.Context) {
	ctx.Set(reflect.ChanOf(reflect.SendDir, c.Sender.Type().Elem()), c.Sender)
	ctx.Set(reflect.ChanOf(reflect.RecvDir, c.Receiver.Type().Elem()), c.Receiver)
}

// Waits for a disconnect message and closes the connection with an appropriate close message.
// The possible messages are:
// TODO this should get more elaborate.
// CloseNormalClosure           = 1000
// CloseGoingAway               = 1001
// CloseProtocolError           = 1002
// CloseUnsupportedData         = 1003
// CloseNoStatusReceived        = 1005
// CloseAbnormalClosure         = 1006
// CloseInvalidFramePayloadData = 1007
// ClosePolicyViolation         = 1008
// CloseMessageTooBig           = 1009
// CloseMandatoryExtension      = 1010
// CloseInternalServerErr       = 1011
// CloseTLSHandshake            = 1015
func waitForDisconnect(c Binding) {
	for {
		select {
		case err := <-c.disconnectChannel():
			if err == io.EOF {
				c.ErrorChannel() <- err
				c.Close(websocket.CloseNormalClosure)
			} else {
				c.Close(websocket.CloseAbnormalClosure)
			}

			return
		case closeCode := <-c.DisconnectChannel():
			c.Close(closeCode)
			return
		}
	}
}

// Creates a new JSON Connection
func newBinding(iFace interface{}, ws *websocket.Conn, o *Options) Binding {
	typ := reflect.TypeOf(iFace)

	if typ.Kind() == reflect.String {
		return &MessageConnection{
			newConnection(ws, o),
			make(chan string, o.SendChannelBuffer),
			make(chan string, o.RecvChannelBuffer),
		}
	}

	return &JSONConnection{
		newConnection(ws, o),
		makeChanOfType(typ, o.SendChannelBuffer),
		makeChanOfType(typ, o.RecvChannelBuffer),
	}
}

// Creates a new Connection
func newConnection(ws *websocket.Conn, o *Options) *Connection {
	return &Connection{
		o,
		ws,
		ws.RemoteAddr(),
		make(chan error, 1),
		make(chan int, 1),
		make(chan bool, 3),
		make(chan error, 1),
		make(chan bool, 1),
		nil,
	}
}

// Creates new default options and assigns any given options
func newOptions(options []*Options) *Options {
	o := Options{
		log.New(os.Stdout, "[sockets] ", 0),
		defaultLogLevel,
		false,
		defaultWriteWait,
		defaultPongWait,
		defaultPingPeriod,
		defaultMaxMessageSize,
		defaultSendChannelBuffer,
		defaultRecvChannelBuffer,
	}

	// when all defaults, return it
	if len(options) == 0 {
		return &o
	}

	// map the given values to the options
	optionsValue := reflect.ValueOf(options[0])
	oValue := reflect.ValueOf(&o)
	numFields := optionsValue.Elem().NumField()

	for i := 0; i < numFields; i++ {
		if value := optionsValue.Elem().Field(i); value.IsValid() && value.CanSet() && isNonEmptyOption(value) {
			oValue.Elem().Field(i).Set(value)
		}
	}

	return &o
}

// Create a chan of the given type as a reflect.Value
func makeChanOfType(typ reflect.Type, chanBuffer int) reflect.Value {
	return reflect.MakeChan(reflect.ChanOf(reflect.BothDir, reflect.PtrTo(typ)), chanBuffer)
}

// Upgrade the connection to a websocket connection
func upgradeRequest(resp http.ResponseWriter, req *http.Request, o *Options) (*websocket.Conn, int, error) {
	if req.Method != "GET" {
		o.log("Method %s is not allowed", LEVEL_WARNING, req.RemoteAddr, req.Method)
		return nil, http.StatusMethodNotAllowed, errors.New("Method not allowed")
	}

	o.log("Request to %s has been allowed for origin %s", LEVEL_DEBUG, req.RemoteAddr, req.Host, req.Header.Get("Origin"))

	ws, err := websocket.Upgrade(resp, req, nil, 1024, 1024)
	if handshakeErr, ok := err.(websocket.HandshakeError); ok {
		o.log("Handshake failed: %s", LEVEL_WARNING, req.RemoteAddr, handshakeErr)
		return nil, http.StatusBadRequest, handshakeErr
	} else if err != nil {
		o.log("Handshake failed: %s", LEVEL_WARNING, req.RemoteAddr, err)
		return nil, http.StatusBadRequest, err
	}

	o.log("Connection established", LEVEL_INFO, req.RemoteAddr)
	return ws, http.StatusOK, nil
}

func isNonEmptyOption(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() != 0
	case reflect.Bool:
		return v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return v.Float() != 0
	case reflect.Interface, reflect.Ptr:
		return !v.IsNil()
	}
	return false
}
