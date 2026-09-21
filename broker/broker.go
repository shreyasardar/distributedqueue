package broker

import (
	"distributedqueue/models"
	"distributedqueue/topic"
	"sync"
)

type Broker struct {
	Topics map[string]*topic.Topic
	NextID int64
	mu     sync.Mutex
}

func NewBroker() *Broker {
	return &Broker{
		Topics: make(map[string]*topic.Topic),
		NextID: 1,
	}
}

func (b *Broker) CreateTopic(name string) {
	if _, exists := b.Topics[name]; exists {
		return
	}

	b.Topics[name] = topic.NewTopic(name)
}

func (b *Broker) GetTopic(name string) (*topic.Topic, bool) {
	t, exists := b.Topics[name]

	if !exists {
		return nil, false
	}

	return t, true
}

func (b *Broker) Publish(topicName string, msg models.Message) bool {
	topic, ok := b.GetTopic(topicName)

	if !ok {
		return false
	}

	b.mu.Lock()

	msg.ID = b.NextID
	b.NextID++

	b.mu.Unlock()

	topic.Queue.Enqueue(msg)

	return true
}

func (b *Broker) Consume(topicName string) *models.Message {
	topic, ok := b.GetTopic(topicName)

	if !ok {
		return nil
	}

	return topic.Queue.Dequeue()
}
