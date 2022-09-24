package maps_test

import (
	"reflect"
	"testing"

	"github.com/pettinz/go-utils/maps"
)

func TestKeys(t *testing.T) {
	t.Run("TestKeys must return an empty slice", func(t *testing.T) {
		m := make(map[string]string)

		want := make([]string, 0)
		got := maps.Keys(m)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("TestKeys must return the filtered slice", func(t *testing.T) {
		m := map[string]string{
			"key1": "value1",
			"key2": "value2",
		}

		want := []string{"key1", "key2"}
		got := maps.Keys(m)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
