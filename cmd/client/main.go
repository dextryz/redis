package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dextryz/redis"
)

func main() {

	c := redis.OpenClient(":8080")
	//defer c.Close()

	c.Set("a", "b")

    time.Sleep(2*time.Second)

	err, v := c.Get("a")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)
}
