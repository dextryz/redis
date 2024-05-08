package redis

import (
	"encoding/gob"
	"fmt"
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

	// Encode the data onto the connection.
	err := gob.NewEncoder(c.network).Encode(m)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("data stored on server", "key", k, "value", v)
}

func (c *client) Get(k string) (any, error) {

	fmt.Println("getting")

	m := Message{
		Fn:  "get",
		Key: k,
	}

	// Request the value from the server.
	err := gob.NewEncoder(c.network).Encode(m)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for server to response
	time.Sleep(1 * time.Second)

	// Retrieve the response from the connection.
	// The result of Value is encode into the message that the server sends back.
	err = gob.NewDecoder(c.network).Decode(&m)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("data retrieved from server", "key", k, "value", m.Value)

	return m.Value, nil
}
