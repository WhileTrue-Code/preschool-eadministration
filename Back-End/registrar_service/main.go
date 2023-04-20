package main

import (
	"registrar_service/startup"
	"registrar_service/startup/config"
)

func main() {

	cfg := config.NewConfig()
	server := startup.NewServer(cfg)
	server.Start()
}
