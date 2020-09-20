package stack

import (
	"testing"
)

func TestStackOperations(t *testing.T) {
	stack := &Stack{}
	stack.New()

	element1 := &Element{1}
	element3 := &Element{3}
	element5 := &Element{5}
	element7 := &Element{7}

	stack.Push(element1)
	stack.Push(element3)
	stack.Push(element5)
	stack.Push(element7)

	if stack.Pop().Value != 7 {
		t.Error("Expected Value to equal 7.")
	}
}
