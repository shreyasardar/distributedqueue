package broker

import (
	"gominimq/models"
	"sync"
	"testing"
)

func TestConcurrentPublish(t *testing.T) {
	dataDir := t.TempDir()
	b := NewBroker(dataDir)

	err := b.CreateTopic("orders")
	if err != nil {
		t.Fatalf("failed to create topic: %v", err)
	}

	topic, ok := b.GetTopic("orders")

	if !ok {
		t.Fatal("orders topic not found")
	}
	defer topic.Log.Close()

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			msg := models.Message{
				Value: "test",
			}

			b.Publish("orders", msg)
		}()
	}

	wg.Wait()

	if topic.Queue.Size() != 100 {
		t.Errorf("expected 100 messages, got %d", topic.Queue.Size())

	}

	ids := make(map[int64]bool)

	for _, msg := range topic.Queue.Messages {
		if ids[msg.ID] {
			t.Errorf("duplicate message ID found: %d", msg.ID)
		}

		ids[msg.ID] = true
	}
	if len(ids) != 100 {
		t.Errorf("expected 100 unique IDs, got %d", len(ids))
	}
}

func TestPublishPersistsMessage(t *testing.T) {

	dataDir := t.TempDir()
	b := NewBroker(dataDir)

	err := b.CreateTopic("orders")
	if err != nil {
		t.Fatalf("failed to create topic: %v", err)
	}

	topic, ok := b.GetTopic("orders")
	if !ok {
		t.Fatal("orders topic not found")
	}
	defer topic.Log.Close()

	msg := models.Message{
		Key:       "user1",
		Value:     "hello",
		Timestamp: 123456,
	}

	success := b.Publish("orders", msg)

	if !success {
		t.Fatal("expected publish to succeed")
	}

	messages, err := topic.Log.ReadAll()
	if err != nil {
		t.Fatalf("failed to read persisted messages: %v", err)
	}

	if len(messages) != 1 {
		t.Fatalf("expected 1 persisted message, got %d", len(messages))
	}

	if messages[0].Value != "hello" {
		t.Errorf("expected value %q, got %q", "hello", messages[0].Value)
	}

	if messages[0].Key != "user1" {
		t.Errorf("expected key %q, got %q", "user1", messages[0].Key)
	}
}
