package redis

import (
	"encoding/gob"
	"log"
	"net"
)

type Message struct {
	Fn    string // get/set
	Key   string // Can never be empty
	Value any    // Empty if Fn is 'get'
}

// Decode what ever data is on the connection into the data structure.
func DecodeMessage(c net.Conn) (*Message, error) {
	var m Message
	err := gob.NewDecoder(c).Decode(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Encode the data onto the connection.
func (m *Message) Encode(c net.Conn) {
	err := gob.NewEncoder(c).Encode(m)
	if err != nil {
		log.Fatal(err)
	}
}
