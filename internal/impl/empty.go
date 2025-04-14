package impl

import (
	"github.com/lvjp/go-streams/function"
)

type EmptyStream[T any] struct{}

func (*EmptyStream[T]) AllMatch(predicate function.Predicate[T]) bool {
	return false
}

func (*EmptyStream[T]) AnyMatch(predicate function.Predicate[T]) bool {
	return false
}

func (*EmptyStream[T]) NoneMatch(predicate function.Predicate[T]) bool {
	return true
}

func (*EmptyStream[T]) FindAny() (T, bool) {
	var zero T
	return zero, false
}

func (*EmptyStream[T]) FindFirst() (T, bool) {
	var zero T
	return zero, false
}

func (*EmptyStream[T]) Min(comparator function.Comparator[T]) (T, bool) {
	var zero T
	return zero, false
}

func (*EmptyStream[T]) Max(comparator function.Comparator[T]) (T, bool) {
	var zero T
	return zero, false
}

func (*EmptyStream[T]) Count() uint64 {
	return 0
}

func (*EmptyStream[T]) ForEach(consumer function.Consumer[T]) {
	// no-op
}

func (*EmptyStream[T]) Reduce(accumulator function.BinaryOperator[T]) (T, bool) {
	var zero T
	return zero, false
}

func (*EmptyStream[T]) ReduceWithIdentity(identity T, accumulator function.BinaryOperator[T]) T {
	return identity
}
