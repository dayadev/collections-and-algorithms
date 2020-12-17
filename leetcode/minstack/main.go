package main

func main() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-1)
	println(stack.GetMin())
	stack.Pop()
	println(stack.Top())
	println(stack.GetMin())

}

//MinStack defines min stack
type MinStack struct {
	elements, min []int
	size          int
}

/** initialize your data structure here. */

//Constructor inits the min stack
func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0), 0}
}

// Push pushes to top of min stack
func (s *MinStack) Push(x int) {
	s.elements = append(s.elements, x)
	if s.size == 0 {
		s.min = append(s.min, x)
	} else {
		min := s.GetMin()
		if x < min {
			s.min = append(s.min, x)
		} else {
			s.min = append(s.min, min)
		}
	}
	s.size++
}

// Pop pops topmost item
func (s *MinStack) Pop() {
	s.size--
	s.min = s.min[:s.size]
	s.elements = s.elements[:s.size]
}

// Top gets top items
func (s *MinStack) Top() int {
	return s.elements[s.size-1]
}

// GetMin gets min
func (s *MinStack) GetMin() int {
	return s.min[s.size-1]
}
