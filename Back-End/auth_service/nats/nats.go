package nats

import (
	"github.com/nats-io/nats.go"
	"log"
	"os"
)

func Conn() *nats.Conn {

	// NATS_URI is: nats://nats:4222 because instead of "localhost" we use name of container "nats"

	conn, err := nats.Connect(os.Getenv("NATS_URI"))
	if err != nil {
		log.Println("NotConnected")
		log.Fatal(err)
	}
	log.Println("Connected")
	return conn
}
