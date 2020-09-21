package list

import (
	"container/list"
	"testing"
)

func TestListOperations(t *testing.T) {
	l := list.New()
	four := l.PushBack(4)
	l.PushBack(5)
	one := l.PushFront(1)

	l.InsertAfter(2, one)
	l.InsertBefore(3, four)

	count := 1
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value != count {
			t.Error("Should equal count")
		}
		count++
	}

	if l.Len() != 5 {
		t.Error("The length should equal 5")
	}
}
