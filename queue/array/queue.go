package queue

import "sync"

// Queue holds an array of items
type Queue struct {
	items []interface{}
	lock  sync.RWMutex
}

// New creates a Queue
func New() *Queue {
	q := &Queue{
		items: []interface{}{},
	}
	return q
}

// IsEmpty returns if the Queue is empty or not
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// EnQueue adds item to end of the Queue
func (q *Queue) EnQueue(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, item)
}

// DeQueue removes the first item from the Queue
func (q *Queue) DeQueue() (interface{}, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.IsEmpty() {
		return nil, false
	}
	index := len(q.items) - 1
	item := (q.items)[0]
	if index == 0 {
		q.items = nil
	} else {
		q.items = (q.items)[1:index] // remove it from the Queue by slicing it off
	}
	return item, true
}
