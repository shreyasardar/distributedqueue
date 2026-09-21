package queue

import (
	"distributedqueue/models"
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
