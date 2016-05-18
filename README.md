# go-queue
Implements a damn insanely simple queue ADT in GO lang.

This project is part of my "learn go" projects.

## Types

* queue.Queue
 
## Functions

* NewQueue() *Queue - Creates a new queue
 
## Methods

* (Queue *) Enqueue(interface{}) - Adds a new element to the queue with the given value.
* (Queue *) Dequeue() interface{} - Gets the next element from the queue.
* (Queue *) Empty() bool - Determines whether the queue is empty or not.
 
## Example

```go
queue := queue.NewQueue()
for i := 0; i < 10; i++ {
  queue.Enqueue(i)
}
for !queue.Empty() {
  fmt.Println(queue.Dequeue().(int));
}
```
