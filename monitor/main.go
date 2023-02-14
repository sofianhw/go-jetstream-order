package main

import (
	"encoding/json"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"

	"github.com/sofianhw/go-jetstream-order/model"
)

func main() {
	// Connect to NATS
	// nc, _ := nats.Connect(nats.DefaultURL)
	nc, _ := nats.Connect("nats://172.17.0.1:4222")
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}
	// Create durable consumer monitor
	js.Subscribe("ORDERS.*", func(msg *nats.Msg) {
		msg.Ack()
		var order model.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("\nmonitor service subscribes from subject:%s\n", msg.Subject)
		log.Printf("OrderID:%d, CustomerID: %s, Status:%s\n\n", order.OrderID, order.CustomerID, order.Status)
	}, nats.Durable("monitor"), nats.ManualAck())

	runtime.Goexit()

}
