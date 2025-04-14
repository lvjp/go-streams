package impl

import (
	"iter"

	"github.com/lvjp/go-streams/function"
)

func NewSeq[T any](seq iter.Seq[T]) *Seq[T] {
	return &Seq[T]{seq: seq}
}

type Seq[T any] struct {
	seq iter.Seq[T]
}

func (s *Seq[T]) AllMatch(predicate function.Predicate[T]) bool {
	for v := range s.seq {
		if !predicate(v) {
			return false
		}
	}

	return true
}

func (s *Seq[T]) AnyMatch(predicate function.Predicate[T]) bool {
	for v := range s.seq {
		if predicate(v) {
			return true
		}
	}

	return false
}

func (s *Seq[T]) NoneMatch(predicate function.Predicate[T]) bool {
	for v := range s.seq {
		if predicate(v) {
			return false
		}
	}

	return true
}

func (s *Seq[T]) FindAny() (T, bool) {
	return s.FindFirst()
}

func (s *Seq[T]) FindFirst() (T, bool) {
	for v := range s.seq {
		return v, true
	}

	var zero T
	return zero, false
}

func (s *Seq[T]) Min(comparator function.Comparator[T]) (T, bool) {
	var found bool
	var v T
	for i := range s.seq {
		if !found {
			v = i
			found = true
			continue
		}

		if comparator(i, v) < 0 {
			v = i
		}
	}

	return v, found
}

func (s *Seq[T]) Max(comparator function.Comparator[T]) (T, bool) {
	var found bool
	var res T
	for v := range s.seq {
		if !found {
			res = v
			found = true
			continue
		}

		if comparator(v, res) > 0 {
			res = v
		}
	}

	return res, found
}

func (s *Seq[T]) Count() uint64 {
	var count uint64
	for range s.seq {
		count++
		if count == 0 {
			panic("count overflow")
		}
	}

	return count
}

func (s *Seq[T]) ForEach(consumer function.Consumer[T]) {
	for v := range s.seq {
		consumer(v)
	}
}

func (s *Seq[T]) Reduce(accumulator function.BinaryOperator[T]) (T, bool) {
	var found bool
	var res T
	for v := range s.seq {
		if !found {
			res = v
			found = true
			continue
		}
		res = accumulator(res, v)
	}

	return res, found
}

func (s *Seq[T]) ReduceWithIdentity(identity T, accumulator function.BinaryOperator[T]) T {
	res := identity

	for v := range s.seq {
		res = accumulator(res, v)
	}

	return res
}
