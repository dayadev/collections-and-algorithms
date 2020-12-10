package main

import (
	"time"

	stack "github.com/dayadev/collections-and-algorithms/stack"
)

func main() {
	s := stack.New()
	go s.Push("Car")
	go s.Push("Bus1")
	go s.Push("Bus2")
	go s.Push("Bus3")
	go s.Push("Bus4")
	go s.Push("Bus5")
	go s.Push("Bus6")
	time.Sleep(10 * time.Second)

	popped, _ := s.Pop()
	println(popped.(string))
}
