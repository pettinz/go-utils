package slices_test

import (
	"reflect"
	"testing"

	"github.com/pettinz/go-utils/slices"
)

func TestFilter(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	t.Run("TestFilter must return an empty slice", func(t *testing.T) {
		want := make([]int, 0)
		got := slices.Filter(numbers, func(_ int, v int) bool {
			return v > 10
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("TestFilter must return the filtered slice", func(t *testing.T) {
		want := []int{0, 1, 2, 3, 4}
		got := slices.Filter(numbers, func(_ int, v int) bool {
			return v < 5
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
