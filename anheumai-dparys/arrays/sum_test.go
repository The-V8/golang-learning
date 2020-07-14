package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("Sum all slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		checkSums(t, got, want)
	})

	t.Run("Sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

	t.Run("Sum without head and tail", func(t *testing.T) {
		got := SumWithoutHeadAndTail([]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{1, 1, 2, 2})
		want := []int{5, 13, 3}
		checkSums(t, got, want)
	})

	t.Run("Sum without head and tail and one empty slice", func(t *testing.T) {
		got := SumWithoutHeadAndTail([]int{1, 2, 3, 4}, []int{}, []int{5, 6, 7, 8}, []int{1, 1, 2, 2})
		want := []int{5, 0, 13, 3}
		checkSums(t, got, want)
	})
}
