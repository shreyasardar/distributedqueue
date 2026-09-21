package api

import (
	"distributedqueue/broker"
	"distributedqueue/models"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Broker *broker.Broker
}

func NewServer(b *broker.Broker) *Server {
	return &Server{
		Broker: b,
	}
}

type PublishRequest struct {
	Topic string `json:"topic"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s *Server) PublishHandler(c *gin.Context) {

	var req PublishRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	msg := models.Message{
		Key:   req.Key,
		Value: req.Value,
	}

	s.Broker.Publish(req.Topic, msg)

	c.JSON(200, gin.H{
		"message": "Message published successfully",
	})
}

func (s *Server) ConsumeHandler(c *gin.Context) {

	topicName := c.Query("topic")

	msg := s.Broker.Consume(topicName)

	if msg == nil {
		c.JSON(404, gin.H{
			"error": "No message available",
		})
		return
	}

	c.JSON(200, msg)
}
