package queue

import (
	"math"
)

type Queue[T any] struct {
	start int
	end   int
	len   int
	data  []T
}

// Returns a new Queue object but will panic if you pass in a negative `initialCapacity`
func MakeQueue[T any](initialCapacity int) Queue[T] {
	return Queue[T]{
		start: 0,
		end:   0,
		data:  make([]T, initialCapacity),
	}
}

func (q *Queue[T]) GrowBy(n int) bool {
	if n < 1 {
		return false
	}

	newData := make([]T, cap(q.data)+n)
	cur := 0
	for i := q.start; i < len(q.data); i += 1 {
		newData[cur] = q.data[i]
		cur += 1
	}
	for i := 0; i < q.end; i += 1 {
		newData[cur] = q.data[i]
		cur += 1
	}

	q.data = newData
	q.start = 0
	q.end = q.len
	return true
}

func (q *Queue[T]) Grow() {
	growBy := math.Ceil(float64(cap(q.data)) * 0.5)
	q.GrowBy(int(growBy))
}

func (q *Queue[T]) Enqueue(item T) bool {
	if q.len == len(q.data) {
		q.Grow()
	}
	q.data[q.end] = item
	q.end = (q.end + 1) % len(q.data)
	q.len += 1
	return true
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.len == 0 {
		var zero T
		return zero, false
	}
	item := q.data[q.start]
	q.start = (q.start + 1) % len(q.data)
	q.len -= 1
	return item, true
}

func (q *Queue[T]) Len() int {
	return q.len
}

func (q *Queue[T]) Front() (T, bool) {
	if q.len == 0 {
		var zero T
		return zero, false
	}
	return q.data[q.start], true
}

func (q *Queue[T]) Back() (T, bool) {
	if q.len == 0 {
		var zero T
		return zero, false
	}
	idx := (q.end + len(q.data) - 1) % len(q.data)
	return q.data[idx], true
}
