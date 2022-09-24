package maps_test

import (
	"testing"

	"github.com/pettinz/go-utils/maps"
)

func TestUnzip(t *testing.T) {
	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	k, v := maps.Unzip(m)

	if !(m[k[0]] == v[0] && m[k[1]] == v[1]) {
		t.Errorf("got keys %v, %v on %v", k, v, m)
	}
}
