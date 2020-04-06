package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"time"
)

func connect() (*nats.Conn, error) {
	nc, err = nats.Connect(
		os.Getenv("NATS_URL"),
		nats.Name(os.Getenv("SERVER_NAME")),
		nats.Timeout(10*time.Second),
		nats.PingInterval(20*time.Second),
		nats.MaxPingsOutstanding(5),
		nats.MaxReconnects(10),
		nats.ReconnectWait(10*time.Second),
		nats.ReconnectBufSize(5*1024*1024))
	if err != nil {
		return nc, err
	}
	log.Println("Connected to NATS Server with URL ", nc.ConnectedUrl())
	return nc, nil
}
