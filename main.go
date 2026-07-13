package main

import (
	"distributedqueue/broker"
	"distributedqueue/models"
	"fmt"
)

func main() {

	b := broker.NewBroker()

	b.CreateTopic("orders")

	msg1 := models.Message{
		ID:        1,
		Key:       "order_101",
		Value:     "Pizza",
		Timestamp: 109,
	}

	msg2 := models.Message{
		ID:        2,
		Key:       "order_102",
		Value:     "Pasta",
		Timestamp: 110,
	}

	b.Publish("orders", msg1)
	b.Publish("orders", msg2)

	fmt.Println(b.Consume("orders"))
	fmt.Println(b.Consume("orders"))
}
