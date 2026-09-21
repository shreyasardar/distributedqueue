package broker

import (
	"distributedqueue/models"
	"sync"
	"testing"
)

func TestConcurrentPublish(t *testing.T) {
	b := NewBroker()

	b.CreateTopic("orders")

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

	topic, ok := b.GetTopic("orders")

	if !ok {
		t.Fatal("orders topic not found")
	}

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
