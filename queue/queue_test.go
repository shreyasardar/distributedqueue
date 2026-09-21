package queue

import (
	"gominimq/models"
	"sync"
	"testing"
)

func TestConcurrentEnqueue(t *testing.T) {
	q := NewQueue()

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			q.Enqueue(models.Message{
				Value: "test",
			})
		}()
	}

	wg.Wait()

	if q.Size() != 100 {
		t.Errorf("expected 100 messages, got %d", q.Size())
	}
}

func TestConcurrentEnqueueDequeue(t *testing.T) {
	q := NewQueue()

	var wg sync.WaitGroup

	// 100 producers
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			q.Enqueue(models.Message{
				Value: "test",
			})
		}()
	}

	wg.Wait()
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			q.Dequeue()
		}()
	}

	wg.Wait()

	if !q.IsEmpty() {
		t.Errorf("expected queue to be empty, got %d messages", q.Size())
	}
}
