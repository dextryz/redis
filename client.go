package redis

import (
	"log"
	"log/slog"
	"net"
	"time"
)

type client struct {
	network net.Conn
}

func OpenClient(port string) *client {

	c, err := net.Dial("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	return &client{
		network: c,
	}
}

func (c *client) Close() {
	err := c.network.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (c *client) Set(k string, v any) {

	m := Message{
		Fn:    "set",
		Key:   k,
		Value: v,
	}

	m.Encode(c.network)

	slog.Info("data sent to server", "key", k, "value", v)
}

func (c *client) Get(k string) (any, error) {

	m := Message{
		Fn:  "get",
		Key: k,
	}

	// Request the value from the server.
	m.Encode(c.network)

	// Wait for server to response
	time.Sleep(1 * time.Second)

	// Retrieve the response from the connection.
	// The result of Value is encode into the message that the server sends back.
	mr, err := DecodeMessage(c.network)
	if err != nil {
		return nil, err
	}

	slog.Info("response from server", "key", k, "value", mr.Value)

	return mr.Value, nil
}
