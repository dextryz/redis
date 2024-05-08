package redis

import (
	"fmt"
	"io"
	"log"
	"net"
)

type srv struct {
	ln    net.Listener
	store *store
}

func NewServer(port string) *srv {

	store := OpenStore("")

	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("XXX")
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

				m, err := DecodeMessage(c)
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Println("AAA")
					log.Fatal(err)
				}

				if m.Fn == "set" {
					s.store.Set(m.Key, m.Value)
				}

				if m.Fn == "get" {
					v, err := s.store.Get(m.Key)
					if err != nil {
						log.Fatal(err)
					}
					m.Value = v
					m.Encode(c)
				}

			}

			c.Close()
		}(conn)
	}
}
