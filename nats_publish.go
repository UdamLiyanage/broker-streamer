package main

import (
	"github.com/nats-io/nats.go"
	"github.com/prometheus/common/log"
	"os"
)

type (
	NATSMessage struct {
		Topic   string
		Payload []byte
	}
)

func publish(payload []byte, topic string) {
	msgPayload := NATSMessage{
		Topic:   topic,
		Payload: payload,
	}
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Error(err)
	}
	defer ec.Close()
	if err := ec.Publish(os.Getenv("NATS_SUBJECT"), &msgPayload); err != nil {
		log.Error(err)
	}
}
