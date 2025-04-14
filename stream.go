package gostream

import (
	"errors"
	"iter"

	"github.com/lvjp/go-streams/function"
	"github.com/lvjp/go-streams/internal/impl"
)

type Stream[T any] interface {
	// Parallel() Stream[T]
	// IsParallel() bool

	// Sequential() Stream[T]
	// Unordered() Stream[T]

	// Peek(consumer Consumer[T]) Stream[T]
	// Filter(predicate Predicate[T]) Stream[T]
	// Limit(maxSize uint64) Stream[T]
	// Skip(n uint64) Stream[T]
	// Sorted() Stream[T]
	// SortedBy(comparator Comparator[T]) Stream[T]
	// Distinct() Stream[T]
	// DistinctBy(comparator Comparator[T]) Stream[T]

	FindAny() (T, bool)
	FindFirst() (T, bool)
	Count() uint64

	ForEach(consumer function.Consumer[T])
	// ForEachUnordered(consumer Consumer[T])
	// ForEachOrdered(consumer Consumer[T])

	AllMatch(predicate function.Predicate[T]) bool
	AnyMatch(predicate function.Predicate[T]) bool
	NoneMatch(predicate function.Predicate[T]) bool
	Min(comparator function.Comparator[T]) (T, bool)
	Max(comparator function.Comparator[T]) (T, bool)

	Reduce(accumulator function.BinaryOperator[T]) (T, bool)
	ReduceWithIdentity(identity T, accumulator function.BinaryOperator[T]) T

	// Parameterized methods are not supported in Go.
	// https://github.com/golang/proposal/blob/master/design/43651-type-parameters.md#no-parameterized-methods
	// FlatMap[U any](mapper function.Function[T,Stream[U]]) Stream[U]
	// Map[U any](mapper function.Function[T,U]) Stream[U]
	// Collect[A,R any](collector Collector[T,A,R]) R
}

func NewEmptyStream[T any]() Stream[T] {
	return &impl.EmptyStream[T]{}
}

func NewStreamOfValues[T any](values ...T) Stream[T] {
	return impl.NewSlice(values)
}

func NewStreamOfSeq[T any](seq iter.Seq[T]) Stream[T] {
	return impl.NewSeq(seq)
}

func NewIterativeStream[T any](seed T) Stream[T] {
	panic(errors.ErrUnsupported)
}
