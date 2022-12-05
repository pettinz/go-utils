package slices_test

import (
	"reflect"
	"testing"

	"github.com/pettinz/go-utils/slices"
)

func TestUnique(t *testing.T) {
	letters := []string{"a", "b", "c", "a", "b", "a"}

	t.Run("TestUnique must return an empty slice if it is passed", func(t *testing.T) {
		want := make([]string, 0)
		got := slices.Unique([]string{})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("TestUnique must return the slice without duplicates", func(t *testing.T) {
		want := []string{"a", "b", "c"}
		got := slices.Unique(letters)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
