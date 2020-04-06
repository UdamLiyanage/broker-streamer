package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"testing"
)

func TestMqttConnection(t *testing.T) {
	opts := mqtt.NewClientOptions().AddBroker(os.Getenv("BROKER_URL"))
	opts.SetUsername(os.Getenv("BROKER_USERNAME"))
	opts.SetPassword(os.Getenv("BROKER_PASSWORD"))
	topic := "/#"

	opts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(topic, 0, func(client mqtt.Client, message mqtt.Message) {
			println(message.Topic())
		}); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
	}
}
