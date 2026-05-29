package main

import (
	"avadnet/backend/internal/api"
	"avadnet/backend/internal/config"
)

func main() {
	cfg := config.Load()
	router := api.NewRouter(cfg)
	api.Run(router, cfg)
}
