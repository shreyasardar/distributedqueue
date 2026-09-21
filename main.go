package main

import (
	"distributedqueue/api"
	"distributedqueue/broker"

	"github.com/gin-gonic/gin"
)

func main() {

	// Create broker
	b := broker.NewBroker()

	b.CreateTopic("orders")

	// Create API server using the broker
	server := api.NewServer(b)

	// Create Gin router
	router := gin.Default()

	// Test endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start server
	router.POST("/publish", server.PublishHandler)

	router.GET("/consume", server.ConsumeHandler)

	router.Run(":8080")
}
