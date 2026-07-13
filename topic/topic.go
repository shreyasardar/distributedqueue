package topic

import (
	"distributedqueue/queue"
)

type Topic struct {
	Name  string
	Queue *queue.Queue
}

func NewTopic(name string) *Topic {
	return &Topic{
		Name:  name,
		Queue: queue.NewQueue(),
	}
}