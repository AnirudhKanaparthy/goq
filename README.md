# goq (GOQueue)
A simple queue implementation in Go using ring buffers.

## Features
1. Simple API
2. Grows as you add elements
3. No wasted memory
4. Fast

## Usage
```go
q := queue.MakeQueue[int](2)

q.Enqueue(23)
q.Enqueue(45)
q.Enqueue(67)

if q.Len() != 3 {
    panic("Please report an issue")
}

item, ok := q.Dequeue()
item, ok = q.Dequeue()
```

## Important
This is not a thread-safe implementation.

## Why?
I was quiet surprised when the standard library of Go didn't provide any Queue implementation. Some people resorted to using the following pattern:
```go
slice = append(slice, item) // Enqueue
item = slice[0] // Back
slice = slice[1:] // Dequeue
```
The problem with this pattern is that we waste a lot of memory, which could be ok if the number of elements is small. But it eats up memory very quickly if there are a lot of elements.

Another method requires the use of a LinkedList. Which can be slower because the memory is not all in a single place.

Ring buffers seemed like the best way of implementing the Queue.