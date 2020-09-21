package linkedlist

import (
	"testing"
)

func TestSingleLinkedList(t *testing.T) {
	list := LinkedList{}
	list.Append(5)
	list.Prepend(2)
	list.Prepend(1)
	if list.FindNode(2).Property != 2 {
		t.Error("Should have found 2")
	}
}
