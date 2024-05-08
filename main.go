package main

import (
	"encoding/gob"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func mustCopy(dst io.Writer, src io.Reader) {

	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	c, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Print the response from the server
	go mustCopy(os.Stdout, c)

	d := map[string]any{
		"key": "hello friend",
	}

	// Encode the data onto the connection.
	// THe data type does not have a io.Write or io.REader interface, so use the gob pkg
	err = gob.NewEncoder(c).Encode(d)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(3 * time.Second)
}
