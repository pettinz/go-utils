package maps_test

import (
	"reflect"
	"testing"

	"github.com/pettinz/go-utils/maps"
)

func TestValues(t *testing.T) {
	t.Run("TestValues must return an empty slice", func(t *testing.T) {
		m := make(map[string]string)

		want := make([]string, 0)
		got := maps.Values(m)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("TestValues must return the filtered slice", func(t *testing.T) {
		m := map[string]string{
			"key1": "value1",
			"key2": "value2",
		}

		want := []string{"value1", "value2"}
		got := maps.Values(m)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
