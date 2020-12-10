package stack

import "sync"

// Stack holds an array of items
type Stack struct {
	items []interface{}
	lock  sync.RWMutex
}

// New creates a stack
func New() *Stack {
	s := &Stack{
		items: []interface{}{},
	}
	return s
}

// IsEmpty returns if the stack is empty or not
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Push pushes item to top of the stack
func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, item)
}

// Pop pops the top most item from the stack
func (s *Stack) Pop() (interface{}, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.IsEmpty() {
		return nil, false
	}
	index := len(s.items) - 1   //fetch the index of top most item in the stack
	item := (s.items)[index]    //fetch the item at the index(top most item in the stack)
	s.items = (s.items)[:index] // remove it from the stack by slicing it off
	return item, true
}
