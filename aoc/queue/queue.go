package queue

type Queue[T any] []T

func (q *Queue[T]) Put(x T) {
	*q = append(*q, x)
}

func (q *Queue[T]) Get() T {
	ret := (*q)[0]
	*q = (*q)[1:]

	return ret
}

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}
