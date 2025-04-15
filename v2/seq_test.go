package gostream

import (
	"fmt"
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCount(t *testing.T) {
	testCases := []struct {
		expected uint64
		input    []any
	}{
		{0, []any{}},
		{1, []any{nil}},
		{2, []any{nil, nil}},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			source := slices.Values(tc.input)
			actual := Count(source)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestReduce(t *testing.T) {
	reducer := func(a, b int) int {
		return a * b
	}

	t.Run("normal", func(t *testing.T) {
		// By the fundamental theorem of arithmetic, every positive integer has a unique prime factorization.
		// Thats why we use only prime numbers.
		source := slices.Values([]int{1, 3, 7, 13})
		expected := 1 * 3 * 7 * 13

		value, ok := Reduce(source, reducer)
		require.True(t, ok)
		require.Equal(t, expected, value)
	})

	t.Run("empty", func(t *testing.T) {
		source := slices.Values([]int{})

		_, ok := Reduce(source, reducer)
		require.False(t, ok)
	})
}

func TestMap(t *testing.T) {
	source := slices.Values([]int{0, 1, 3, 5})
	expected := []string{"0b0", "0b1", "0b11", "0b101"}

	mapper := func(v int) string {
		return fmt.Sprintf("0b%b", v)
	}

	actual := Map(source, mapper)
	require.Equal(t, expected, slices.Collect(actual))
}

func TestFlatMap(t *testing.T) {
	source := slices.Values([]int{0, 1, 3, 5})
	expected := []rune{
		// 0 is empty
		'0', 'b', '1', // 1
		'0', 'b', '1', '1', // 3
		'0', 'b', '1', '0', '1', // 5
	}

	mapper := func(v int) iter.Seq[rune] {
		if v == 0 {
			return slices.Values([]rune{})
		}

		return slices.Values([]rune(fmt.Sprintf("0b%b", v)))
	}

	actual := FlatMap(source, mapper)
	require.Equal(t, expected, slices.Collect(actual))
}
