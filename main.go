package main

import (
	"fmt"

	"github.com/AnirudhKanaparthy/goq/queue"
)

func main() {
	q := queue.MakeQueue[int](2)

	fmt.Println(q.Enqueue(23))
	fmt.Println(q.Enqueue(45))
	fmt.Println(q.Enqueue(67))
	fmt.Println(q.Enqueue(75))
	fmt.Println(q.Enqueue(99))

	fmt.Println(q.Front())
	fmt.Println(q.Back())

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	fmt.Println(q.Front())
	fmt.Println(q.Back())

	fmt.Println(q.Enqueue(173))
	fmt.Println(q.Enqueue(954))
	fmt.Println(q.Enqueue(998))

	fmt.Println(q.Front())
	fmt.Println(q.Back())

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	fmt.Println(q.Front())
	fmt.Println(q.Back())

	fmt.Println(q.Enqueue(87))
	fmt.Println(q.Enqueue(90))
	fmt.Println(q.Enqueue(11))
	fmt.Println(q.Enqueue(99))
	fmt.Println(q.Enqueue(65))
	fmt.Println(q.Enqueue(455))
	fmt.Println(q.Enqueue(2333))
	fmt.Println(q.Enqueue(664))

	fmt.Println(q.Front())
	fmt.Println(q.Back())

	fmt.Printf("Queue: %v\n", q)
}
