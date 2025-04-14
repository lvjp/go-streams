package impl

import (
	"slices"

	"github.com/lvjp/go-streams/function"
)

func NewSlice[T any](data []T) *Slice[T] {
	return &Slice[T]{data: data}
}

type Slice[T any] struct {
	data []T
}

func (s *Slice[T]) AllMatch(predicate function.Predicate[T]) bool {
	for _, v := range s.data {
		if !predicate(v) {
			return false
		}
	}

	return true
}

func (s *Slice[T]) AnyMatch(predicate function.Predicate[T]) bool {
	return slices.ContainsFunc(s.data, predicate)
}

func (s *Slice[T]) NoneMatch(predicate function.Predicate[T]) bool {
	return s.AnyMatch(predicate.Negate())
}

func (s *Slice[T]) FindAny() (T, bool) {
	return s.FindFirst()
}

func (s *Slice[T]) FindFirst() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}

	return s.data[0], true
}

func (s *Slice[T]) Min(comparator function.Comparator[T]) (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}

	return slices.MinFunc(s.data, comparator), true
}

func (s *Slice[T]) Max(comparator function.Comparator[T]) (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}

	return slices.MaxFunc(s.data, comparator), true
}

func (s *Slice[T]) Count() uint64 {
	return uint64(len(s.data))
}

func (s *Slice[T]) ForEach(consumer function.Consumer[T]) {
	for _, v := range s.data {
		consumer(v)
	}
}

func (s *Slice[T]) Reduce(accumulator function.BinaryOperator[T]) (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}

	res := s.data[0]
	for i := 1; i < len(s.data); i++ {
		res = accumulator(res, s.data[i])
	}

	return res, true
}

func (s *Slice[T]) ReduceWithIdentity(identity T, accumulator function.BinaryOperator[T]) T {
	res := identity

	for _, v := range s.data {
		res = accumulator(res, v)
	}

	return res
}
