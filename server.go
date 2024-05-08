package redis

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

type srv struct {
	ln    net.Listener
	store *Store
}

func NewServer(port string) *srv {

	store := OpenStore()

	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	return &srv{
		ln:    ln,
		store: store,
	}
}

func (s srv) Run() {

	for {

		conn, err := s.ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {

			// Read values from connection until EOF
			for {

				m := Message{}

				// Decode what ever data is on the connection into the data structure.
				err := gob.NewDecoder(c).Decode(&m)
				if err != nil {
					log.Fatal(err)
				}

				if m.Fn == "set" {
					fmt.Println("has to Set")
					s.store.Set(m.Key, m.Value)
				}

				if m.Fn == "get" {
					fmt.Println("has to Get")
					err, v := s.store.Get(m.Key)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(v)
				}

			}

			//c.Close()
		}(conn)
	}
}
