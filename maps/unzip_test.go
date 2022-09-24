package maps_test

import (
	"reflect"
	"testing"

	"github.com/pettinz/go-utils/maps"
)

func TestUnzip(t *testing.T) {
	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	want_k := []string{"key1", "key2"}
	want_v := []string{"value1", "value2"}
	got_k, got_v := maps.Unzip(m)

	if !reflect.DeepEqual(got_k, want_k) {
		t.Errorf("got keys %v, want keys %v", got_k, want_k)
	}
	if !reflect.DeepEqual(got_v, want_v) {
		t.Errorf("got keys %v, want keys %v", got_v, want_v)
	}
}
