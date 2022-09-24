package slices_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/pettinz/go-utils/slices"
)

func TestMap(t *testing.T) {
	t.Run("TestMap must return an empty slice", func(t *testing.T) {
		numbers := []int{}

		want := []string{}
		got := slices.Map(numbers, func(_ int, v int) string {
			return strconv.Itoa(v)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("TestMap must return a slice", func(t *testing.T) {
		numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		want := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		got := slices.Map(numbers, func(_ int, v int) string {
			return strconv.Itoa(v)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
