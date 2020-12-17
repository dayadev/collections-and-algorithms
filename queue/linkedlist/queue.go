package queue

var first, last *Queue

// Queue holds item and ref to next time in queue
type Queue struct {
	item interface{}
	next *Queue
}

//DeQueue removes the first item in the queue
func (q *Queue) DeQueue() interface{} {
	var item = first.item
	first = first.next
	return item
}

// EnQueue adds an item to the end of the queue
func (q *Queue) EnQueue(item interface{}) {
	var oldLast = last
	last = &Queue{}
	last.item = &item
	last.next = nil
	if q.IsEmpty() {
		first = last
	} else {
		oldLast.next = last
	}
}

// IsEmpty checks if queue is empty or not
func (q *Queue) IsEmpty() bool {
	return first == nil
}
