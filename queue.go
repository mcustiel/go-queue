package queue

import (
	"container/list"
)

type Queue struct {
	internalList *list.List
}

func NewQueue() *Queue {
	queue := new(Queue)
	queue.internalList = list.New()
	return queue
}

func (queue *Queue) Enqueue(value interface{}) {
	queue.internalList.PushBack(value)
}

func (queue *Queue) Dequeue() interface{} {
	element := queue.internalList.Front()
	queue.internalList.Remove(element)
	return element.Value
}

func (queue *Queue) Empty() bool {
	return queue.internalList.Len() == 0
}
