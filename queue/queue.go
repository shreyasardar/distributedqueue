package queue

import (
	"fmt"
	"gominimq/models"
	"sync"
)

type Queue struct {
	Messages []models.Message
	mu       sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		Messages: []models.Message{},
	}
}

func (q *Queue) Enqueue(msg models.Message) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.Messages = append(q.Messages, msg)
}

func (q *Queue) Dequeue() *models.Message {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.Messages) == 0 {
		return nil
	}

	first := q.Messages[0]
	q.Messages = q.Messages[1:]

	return &first
}

func (q *Queue) Peek() *models.Message {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.Messages) == 0 {
		return nil
	}

	return &q.Messages[0]
}

func (q *Queue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.Messages)
}

func (q *Queue) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.Messages) == 0
}

func (q *Queue) Print() {
	q.mu.Lock()
	defer q.mu.Unlock()

	fmt.Println("Queue Messages:")
	fmt.Println("----------------")

	for _, msg := range q.Messages {
		fmt.Println(msg)
	}
}
