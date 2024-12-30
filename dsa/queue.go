package main

type Queue[T comparable] struct {
	vals []T
}

func (q *Queue[T]) Enqueue(val T) {
	q.vals = append(q.vals, val)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.vals) == 0 {
		var zero T

		return zero, false
	}

	first := q.vals[0]
	q.vals = q.vals[1:]

	return first, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.vals) == 0
}

func (q *Queue[T]) Length() int {
	return len(q.vals)
}

func (q *Queue[T]) Contains(val T) bool {
	for _, v := range q.vals {
		if v == val {
			return true
		}
	}

	return false
}
