package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"", ""},
		{"Привет", "тевирП"},
	}

	for _, c := range cases {
		got := rev(c.in)
		if got != c.want {
			t.Errorf("reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
