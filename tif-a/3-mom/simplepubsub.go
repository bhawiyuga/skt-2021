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
	fmt.Println("Terima message dari broker : ", fmt.Sprintf("%s", msg.Payload()), " dengan topik ", msg.Topic())
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
	if args == "pub" {
		pub(client)
	} else if args == "sub" {
		sub(client)
		// Block program agar tidak exit setelah subscribe berhasil
		<-(chan int)(nil)
	}
}

func sub(client mqtt.Client) {
	// Variabel topik message
	topic := "/hello"
	// Subscribe dengan topik tertentu dan QoS Level 1
	token := client.Subscribe(topic, 1, nil)
	fmt.Println("Berhasil subscribe")
	// Menunggu susbcribe berhasil
	token.Wait()
}

func pub(client mqtt.Client) {
	topic := "/hello"
	message := "Selamat pagi"
	client.Publish(topic, 1, false, message)
}
