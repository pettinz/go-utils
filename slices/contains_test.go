package slices_test

import (
	"testing"

	"github.com/pettinz/go-utils/slices"
)

func TestContains(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	t.Run("TestContains must be true", func(t *testing.T) {
		want := true
		got := slices.Contains(numbers, 2)

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("TestContains must be false", func(t *testing.T) {
		want := false
		got := slices.Contains(numbers, 10)

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
}
