package main

import "healthcare_service/startup"
import "healthcare_service/startup/config"

func main() {
	cfg := config.NewConfig()
	server := startup.NewServer(cfg)
	server.Start()
}
