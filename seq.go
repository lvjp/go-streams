package gostream

import (
	"iter"

	"github.com/lvjp/go-streams/function"
)

func Concat[T any](sources ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, source := range sources {
			for t := range source {
				if !yield(t) {
					break
				}
			}
		}
	}
}

func Count[T any](source iter.Seq[T]) uint64 {
	var count uint64
	for range source {
		count++
		if count == 0 {
			panic("count overflow")
		}
	}
	return count
}

func ForEach[T any](source iter.Seq[T], consumer function.Consumer[T]) {
	for v := range source {
		consumer(v)
	}
}

func AllMatch[T any](source iter.Seq[T], predicate function.Predicate[T]) bool {
	for v := range source {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func AnyMatch[T any](source iter.Seq[T], predicate function.Predicate[T]) bool {
	for v := range source {
		if predicate(v) {
			return true
		}
	}
	return false
}

func NoneMatch[T any](source iter.Seq[T], predicate function.Predicate[T]) bool {
	for v := range source {
		if predicate(v) {
			return false
		}
	}
	return true
}

func FindAny[T any](source iter.Seq[T]) (T, bool) {
	return FindFirst(source)
}

func FindFirst[T any](source iter.Seq[T]) (T, bool) {
	for v := range source {
		return v, true
	}

	var zeroValue T
	return zeroValue, false
}

func Max[T any](source iter.Seq[T], comparator function.Comparator[T]) (T, bool) {
	var maxValue T
	var found bool
	for v := range source {
		if !found || comparator(v, maxValue) > 0 {
			maxValue = v
			found = true
		}
	}
	return maxValue, found
}

func Min[T any](source iter.Seq[T], comparator function.Comparator[T]) (T, bool) {
	var minValue T
	var found bool
	for v := range source {
		if !found || comparator(v, minValue) < 0 {
			minValue = v
			found = true
		}
	}
	return minValue, found
}

func Reduce[T any](source iter.Seq[T], accumulator function.BinaryOperator[T]) (T, bool) {
	var result T
	var found bool
	for v := range source {
		if !found {
			result = v
			found = true
		} else {
			result = accumulator(result, v)
		}
	}
	return result, found
}

func ReduceWithIdentity[T any](
	source iter.Seq[T],
	identity T,
	accumulator function.BinaryOperator[T],
) T {
	result := identity
	for v := range source {
		result = accumulator(result, v)
	}
	return result
}

func Map[T, U any](source iter.Seq[T], mapper function.Function[T, U]) iter.Seq[U] {
	return func(yield func(U) bool) {
		for t := range source {
			if !yield(mapper(t)) {
				break
			}
		}
	}
}

func FlatMap[T, U any](source iter.Seq[T], mapper function.Function[T, iter.Seq[U]]) iter.Seq[U] {
	return func(yield func(U) bool) {
		for t := range source {
			seqU := mapper(t)
			for u := range seqU {
				if !yield(u) {
					break
				}
			}
		}
	}
}

func Filter[T any](source iter.Seq[T], predicate function.Predicate[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for t := range source {
			if predicate(t) && !yield(t) {
				break
			}
		}
	}
}

func Limit[T any](source iter.Seq[T], maxSize uint64) iter.Seq[T] {
	return func(yield func(T) bool) {
		var count uint64
		for t := range source {
			if count >= maxSize {
				break
			}
			if !yield(t) {
				break
			}
			count++
		}
	}
}

func TakeWhile[T any](source iter.Seq[T], predicate function.Predicate[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for t := range source {
			if !predicate(t) || !yield(t) {
				break
			}
		}
	}
}

func DropWhile[T any](source iter.Seq[T], predicate function.Predicate[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		drop := true
		for t := range source {
			if drop {
				if predicate(t) {
					continue
				}
				drop = false
			}

			if !yield(t) {
				break
			}
		}
	}
}

func Peek[T any](source iter.Seq[T], consumer function.Consumer[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for t := range source {
			consumer(t)
			if !yield(t) {
				break
			}
		}
	}
}

func Skip[T any](source iter.Seq[T], n uint64) iter.Seq[T] {
	return func(yield func(T) bool) {
		var count uint64
		for t := range source {
			if count < n {
				count++
				continue
			}
			if !yield(t) {
				break
			}
		}
	}
}
