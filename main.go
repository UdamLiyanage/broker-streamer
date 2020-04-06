package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	ec  *nats.EncodedConn
	err error
)

func init() {
	ec, err = connect()
	if err != nil {
		panic(err)
	}
}

func onMessageReceived(_ mqtt.Client, msg mqtt.Message) {
	publish(msg.Payload(), msg.Topic())
	println(string(msg.Payload()))
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	opts := setClientOptions()
	topic := "/#"

	opts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		log.Println("Connected to MQTT Broker")
	}
	<-c
}

func setClientOptions() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions().AddBroker(os.Getenv("BROKER_URL"))
	opts.SetUsername(os.Getenv("BROKER_USERNAME"))
	opts.SetPassword(os.Getenv("BROKER_PASSWORD"))
	opts.SetDefaultPublishHandler(onMessageReceived)
	return opts
}
