package queue

import (
	"testing"
)

func TestQueueOperations(t *testing.T) {
	var queue Queue
	queue.New(5)

	queue.Enqueue(&Node{
		Value: 2,
		Name:  "Pizza",
	})

	queue.Enqueue(&Node{
		Value: 3,
		Name:  "Wings",
	})

	queue.Enqueue(&Node{
		Value: 4,
		Name:  "Coke",
	})

	if queue.Dequeue().Name != "Pizza" {
		t.Error("First item in queue should be Pizza")
	}

	if queue.Dequeue().Name != "Wings" {
		t.Error("where is my wings man!")
	}
}
