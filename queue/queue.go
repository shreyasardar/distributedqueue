package queue

import (
	"distributedqueue/models"
	"fmt"
)

type Queue struct {
	Messages []models.Message
}

func NewQueue() *Queue {
	return &Queue{
		Messages: []models.Message{},
	}
}

func (q *Queue) Enqueue(msg models.Message) {
	q.Messages = append(q.Messages, msg)
}

func (q *Queue) Dequeue() *models.Message {
	if len(q.Messages) == 0 {
		return nil
	}

	first := q.Messages[0]
	q.Messages = q.Messages[1:]

	return &first
}

func (q *Queue) Peek() *models.Message {
	if len(q.Messages) == 0 {
		return nil
	}

	return &q.Messages[0]
}

func (q *Queue) Size() int {
	return len(q.Messages)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Messages) == 0
}

func (q *Queue) Print() {
	fmt.Println("Queue Messages:")
	fmt.Println("----------------")

	for _, msg := range q.Messages {
		fmt.Println(msg)
	}
}