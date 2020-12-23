package stack

// Stack ...
type Stack struct {
	items []int
}

// Constructor ...
func Constructor() Stack {
	return Stack{
		items: []int{},
	}
}

// Push ...
func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}

// Pop ...
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		panic("Empty Stack")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

// IsEmpty ...
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Peek ...
func (s *Stack) Peek() int {
	return s.items[len(s.items)-1]
}
