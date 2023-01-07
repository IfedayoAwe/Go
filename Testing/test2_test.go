package main

import "testing"

// Table driven tests

type value struct {
	S string
	V float64
}

func (v value) String() string {
	return v.S
}

func TestValue(t *testing.T) {
	table := []struct {
		v float64
		s string
	}{
		{1, "1"},
		{2, "2"},
	}

	for _, tt := range table {
		v := value{V: tt.v, S: tt.s}

		if s := v.String(); s != tt.s {
			t.Errorf("%v: wanted %s, got %s", tt.v, tt.s, s)
		}

	}
}
