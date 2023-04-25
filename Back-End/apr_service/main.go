package main

import "apr_service/startup"

func main() {
	server := startup.NewServer()
	server.Start()
}
