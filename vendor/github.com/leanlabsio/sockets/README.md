sockets
=======

Middlware sockets provides WebSockets channels binding for [Macaron](https://github.com/go-macaron/macaron).

[API Reference](https://gowalker.org/github.com/go-macaron/sockets)

### Installation

	go get github.com/go-macaron/sockets

## Description

Package sockets makes it fun to use websockets with Macaron. Its aim is to provide an easy to use interface for socket handling which makes it possible to implement socket messaging with just one `select` statement listening to different channels.

#### JSON

`sockets.JSON` is a simple middleware that organizes websockets messages into any struct type you may wish for.

#### Messages

`sockets.Messages` is a simple middleware that organizes websockets messages into string channels.

## Usage

This package essentially provides a binding of websockets to channels, which you can use as in the following, contrived example:

```go
m.Get("/", sockets.JSON(Message{}), func(receiver <-chan *Message, sender chan<- *Message, done <-chan bool, disconnect chan<- int, errorChannel <-chan error) {
	ticker = time.After(30 * time.Minute)
	for {
		select {
		case msg := <-receiver:
			// here we simply echo the received message to the sender for demonstration purposes
			// In your app, collect the senders of different clients and do something useful with them
			sender <- msg
		case <- ticker:
			// This will close the connection after 30 minutes no matter what
			// To demonstrate use of the disconnect channel
			// You can use close codes according to RFC 6455
			disconnect <- websocket.CloseNormalClosure
		case <-client.done:
			// the client disconnected, so you should return / break if the done channel gets sent a message
			return
		case err := <-errorChannel:
			// Uh oh, we received an error. This will happen before a close if the client did not disconnect regularly.
			// Maybe useful if you want to store statistics
		}
	}
})
```

For a simple string messaging connection with string channels, use `sockets.Message()`.

## Options

You can configure the options for sockets by passing in `sockets.Options` as the second argument to `sockets.JSON` or as the first argument to `sockets.Message`. Use it to configure the following options (defaults are used here):

```go
&sockets.Options{
	// The logger to use for socket logging
	Logger: log.New(os.Stdout, "[sockets] ", 0), // *log.Logger

	// The LogLevel for socket logging, possible values:
	// sockets.LEVEL_ERROR (0)
	// sockets.LEVEL_WARNING (1)
	// sockets.LEVEL_INFO (2)
	// sockets.LEVEL_DEBUG (3)
	
	LogLevel: sockets.LEVEL_INFO, 
	
	// Set to true if you want to skip logging
	SkipLogging: false // bool

	// The time to wait between writes before timing out the connection
	// When this is a zero value time instance, write will never time out
	WriteWait: 60 * time.Second, // time.Duration

	// The time to wait at maximum between receiving pings from the client.
	PongWait: 60 * time.Second, // time.Duration

	// The time to wait between sending pings to the client
	// Attention, this does have to be shorter than PongWait and WriteWait
	// unless you want connections to constantly time out.
	PingPeriod: (60 * time.Second * 8 / 10), // time.Duration

	// The maximum messages size for receiving and sending in bytes
	// Messages bigger than this will lead to panic and disconnect
	MaxMessageSize 65536 // int64

	// The send channel buffer
	// How many messages can be asynchronously held before the channel blocks
	SendChannelBuffer 10 // int

	// The receiving channel buffer
	// How many messages can be asynchronously held before the channel blocks
	RecvChannelBuffer 10 // int
}
```

## Rundown

Since it augments the level of websocket handling, it may prove useful if you know how this package does the channel mapping. Here's the rundown for a connection lifetime:

1. Request Method (must be GET) and origin (can not be cross-origin) are tested. If any of these conditions fail, a HTTP status is returned accordingly and the next handler is ignored.

2. The request is upgraded to a websocket connection. If this fails, `http.StatusBadRequest` is returned and the next handler is ignored.

3. The connection and its channels are created according to given options and mapped for dependency injection in Macaron.

4. Three goroutines are started: One is waiting on the websocket for messages, another is waiting for messages from the next handler and occasionally pinging the client. The third is waiting for disconnection coming from both sides. All these goroutines are closed with the websocket, which is also why you should not try to send messages to a client after an error occurred.

5. On the event of a disconnection sent either by the next handler or via the websocket or when an error occurs, the connection is closed with an appropriate or a given closing message.

## `gorilla` vs `go.net` websockets

The gorilla websocket package is a brilliant implementation of RFC 6455 compliant websockets which has [these advantages over the go.net implementation.](https://github.com/gorilla/websocket#protocol-compliance). 

## Credits

This package is forked from [beatrichartz/martini-sockets](https://github.com/beatrichartz/martini-sockets) with modifications.

## License

This project is under Apache v2 License. See the [LICENSE](LICENSE) file for the full license text.
