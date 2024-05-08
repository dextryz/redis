package main

import (
	"github.com/dextryz/redis"
)

func main() {

	s := redis.NewServer(":8080")
	s.Run()
}
