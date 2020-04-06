package main

import (
	"log"
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
	if err := ec.Publish(os.Getenv("NATS_SUBJECT"), &msgPayload); err != nil {
		log.Println(err)
	}
}
