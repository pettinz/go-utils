package maps_test

import (
	"reflect"
	"testing"

	"github.com/pettinz/go-utils/maps"
)

func TestZip(t *testing.T) {
	k := []string{"key1", "key2"}
	v := []string{"value1", "value2"}

	want := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	got := maps.Zip(k, v)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
