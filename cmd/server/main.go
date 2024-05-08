package main

import (
	"log/slog"

	"github.com/dextryz/redis"
)

func main() {

	slog.Info("starting server")

	s := redis.NewServer(":8080")
	s.Run()
}
