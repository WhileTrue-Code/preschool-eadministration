package main

import "croso_service/startup"

func main() {
	server := startup.NewServer()
	server.Start()
}
