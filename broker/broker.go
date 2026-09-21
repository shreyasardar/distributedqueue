package broker

import (
	"gominimq/models"
	"gominimq/topic"
	"path/filepath"
	"sync"
)

type Broker struct {
	Topics  map[string]*topic.Topic
	NextID  int64
	DataDir string
	mu      sync.Mutex
}

func NewBroker(dataDir string) *Broker {
	return &Broker{
		Topics:  make(map[string]*topic.Topic),
		NextID:  1,
		DataDir: dataDir,
	}
}

func (b *Broker) CreateTopic(name string) error {
	if _, exists := b.Topics[name]; exists {
		return nil
	}

	logFilename := filepath.Join(b.DataDir, name+".log")

	t, err := topic.NewTopic(name, logFilename)
	if err != nil {
		return err
	}

	b.Topics[name] = t

	return nil
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

	err := topic.Log.Append(msg)
	if err != nil {
		return false
	}

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
