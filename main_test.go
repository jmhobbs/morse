package main

import "testing"

func Test_morseToString(t *testing.T) {
	cases := []struct {
		input    uint16
		expected string
	}{
		{0b11, "-"},
		{0b01, "."},
		{0, ""},
		{0b110111, "-.-"},
	}

	for _, c := range cases {
		actual := morseToString(c.input)
		if actual != c.expected {
			t.Errorf("got %v\nwant %v", actual, c.expected)
		}
	}
}
