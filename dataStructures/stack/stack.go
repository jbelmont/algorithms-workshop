package stack

import (
	"strconv"
)

// Element contains stack elements
type Element struct {
	Value int
}

func (e *Element) String() string {
	return strconv.Itoa(e.Value)
}

// Stack is a basic Last In First Out (LIFO) data structure
// the resizes as needed
type Stack struct {
	Elements []*Element
	Count    int
}

// New returns a new stack
func (s *Stack) New() {
	s.Elements = make([]*Element, 0)
}

// Push adds an element to the stack
func (s *Stack) Push(e *Element) {
	s.Elements = append(
		s.Elements[:s.Count],
		e,
	)
	s.Count = s.Count + 1
}

// Pop removes and returns the most recently added element in the stack
func (s *Stack) Pop() *Element {
	if s.Count == 0 {
		return nil
	}

	length := len(s.Elements)
	var element *Element = s.Elements[length-1]
	if length > 1 {
		s.Elements = s.Elements[:length-1]
	} else {
		s.Elements = s.Elements[0:]
	}
	s.Count = len(s.Elements)
	return element
}
