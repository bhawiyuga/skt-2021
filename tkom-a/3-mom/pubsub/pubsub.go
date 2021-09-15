package main

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

func main() {
	var broker = "localhost"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.OnConnect = connectHandler
	opts.SetDefaultPublishHandler(messagePubHandler)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	arg := os.Args[1]
	if arg == "pub" {
		pub(client)
	} else if arg == "sub" {
		sub(client)
		//Block program by receiving from a nil channel
		<-(chan int)(nil)
	}
}

func sub(client mqtt.Client) {
	topic := "/hello"
	token := client.Subscribe(topic, 1, nil)

	fmt.Printf("Subscribed to topic: %s", topic)

	token.Wait()
}

func pub(client mqtt.Client) {
	topic := "/hello"
	message := "Selamat pagi"
	client.Publish(topic, 0, false, message)
}
