package broker

import (
	"distributedqueue/models"
	"distributedqueue/topic"
)

type Broker struct {
	Topics map[string]*topic.Topic
}

func NewBroker() *Broker {
	return &Broker{
		Topics: make(map[string]*topic.Topic),
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

func (b *Broker) Publish(topicName string, msg models.Message) {
	topic, ok := b.GetTopic(topicName)

	if !ok {
		return
	}

	topic.Queue.Enqueue(msg)
}

func (b *Broker) Consume(topicName string) *models.Message {
	topic, ok := b.GetTopic(topicName)

	if !ok {
		return nil
	}

	return topic.Queue.Dequeue()
}