package main

import (
	array_queue "github.com/dayadev/collections-and-algorithms/queue/array"
)

func main() {
	q := array_queue.New()

	_, ok := q.DeQueue()
	println(ok)
	q.EnQueue("Daya")
	dequeued2, ok := q.DeQueue()
	println(ok)
	println(dequeued2.(string))
	q.EnQueue("Daya2")
	dequeued2, ok = q.DeQueue()
	println(ok)
	println(dequeued2.(string))
}
