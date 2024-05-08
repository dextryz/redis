package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

type store struct {
	data map[string]any
}

func main() {

	s := store{
		data: make(map[string]any),
	}

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {

		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {

			// Decode what ever data is on the connection into the data structure.
			err := gob.NewDecoder(c).Decode(&s.data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(s.data)

			// WEite a response message to the client.
			_, err = c.Write([]byte("done"))
			if err != nil {
				log.Fatal(err)
			}

			c.Close()
		}(conn)
	}
}
