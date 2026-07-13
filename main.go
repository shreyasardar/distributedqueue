package main

import (
	"distributedqueue/models"
	"fmt"
)

func main() {
	msg := models.Message{
		ID:        1,
		Key:       "order_101",
		Value:     "Pizza",
		Timestamp: 1720953000,
	}

	fmt.Println(msg)
}
