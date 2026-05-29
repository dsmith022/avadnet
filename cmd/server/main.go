package main

import (
	"avadnet/internal/config"
	"avadnet/internal/server"
)

func main() {
	cfg := config.Load()
	router := server.NewRouter(cfg)
	server.Run(router, cfg)
}
