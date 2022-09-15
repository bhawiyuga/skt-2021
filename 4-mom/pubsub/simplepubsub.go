package main

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Berhasil terhubung ke broker")
}

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("Terima message dari broker : ", string(msg.Payload()), " dengan topik ", msg.Topic())
	//fmt.Println("Terima message")
}

func main() {
	// Deklarasikan opsi untuk koneksi dari pub/sub ke broker
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://127.0.0.1:1883")
	// Deklarasikan callback function untuk handling connection
	opts.OnConnect = connectHandler
	// Deklarasikan callback function untuk handle message masuk
	opts.SetDefaultPublishHandler(messageHandler)
	// Kirim permintaan koneksi MQTT ke broker
	client := mqtt.NewClient(opts)
	// Koneksikan dari client ke broker
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		fmt.Println("Terdapat error koneksi : ", token.Error())
	}

	// Tangkap input dari user untuk menentukan program dijalankan sebagai publisher atau susbcriber
	args := os.Args[1]
	message := "Selamat pagi"
	topic := "/hello"

	if len(os.Args) > 2 {
		topic = os.Args[2]
	}
	if len(os.Args) > 3 {
		message = os.Args[3]
	}
	if args == "pub" {
		pub(client, topic, message)
	} else if args == "sub" {
		sub(client, topic)
		// Block program agar tidak exit setelah subscribe berhasil
		<-(chan int)(nil)
	}
}

func sub(client mqtt.Client, topic string) {
	// Variabel topik message
	// Subscribe dengan topik tertentu dan QoS Level 1
	token := client.Subscribe(topic, 1, nil)
	fmt.Println("Berhasil subscribe :", topic)
	// Menunggu susbcribe berhasil
	token.Wait()
}

func pub(client mqtt.Client, topic string, message string) {
	fmt.Println("Publish message :", message, "topic :", topic)
	token := client.Publish(topic, 1, false, message)
	token.Wait()
}
