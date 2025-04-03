package main

type Num interface {
	int | int8 | int16 | int32 | int64
}

type CircularQueue[T Num] struct {
	values      []T
	front, back int
}

func NewCircularQueue[T Num](size T) CircularQueue[T] {
	v := make([]T, size)
	return CircularQueue[T]{
		values: v,
		front:  -1,
		back:   -1,
	}
}

func (q *CircularQueue[T]) Push(value T) bool {
	if q.Full() {
		return false
	}

	if q.back < len(q.values)-1 {
		q.back = q.back + 1
	} else {
		q.back = 0
	}
	q.values[q.back] = value
	if q.front == -1 {
		q.front = 0
	}

	return true
}

func (q *CircularQueue[T]) Pop() bool {
	if q.Empty() {
		return false
	}
	if q.back == q.front {
		q.back = -1
		q.front = -1
	} else if q.front < len(q.values)-1 {
		q.front = q.front + 1
	} else {
		q.front = 0
	}

	return true
}

func (q *CircularQueue[T]) Front() T {
	if q.front == -1 {
		return -1
	}

	return q.values[q.front]
}

func (q *CircularQueue[T]) Back() T {
	if q.back == -1 {
		return -1
	}

	return q.values[q.back]
}

func (q *CircularQueue[T]) Empty() bool {
	if q.front == -1 {
		return true
	}

	return false
}

func (q *CircularQueue[T]) Full() bool {
	if (q.back == len(q.values)-1 && q.front == 0) ||
		q.front-1 == q.back {
		return true
	}

	return false
}
