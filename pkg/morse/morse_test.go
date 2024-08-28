package morse_test

import (
	"testing"

	"github.com/jmhobbs/morse/pkg/morse"
)

func Test_Morse_String(t *testing.T) {
	cases := []struct {
		input    morse.Morse
		expected string
	}{
		{0b11, "-"},
		{0b01, "."},
		{0, ""},
		{0b110111, "-.-"},
	}

	for _, c := range cases {
		actual := c.input.String()
		if actual != c.expected {
			t.Errorf("got %v\nwant %v", actual, c.expected)
		}
	}
}

func Test_Morse_Rune(t *testing.T) {
	cases := []struct {
		input    morse.Morse
		expected rune
	}{
		{0b01, 'E'},
		{0b010101, 'S'},
		{0b111111, 'O'},
		{0b1111111111, '0'},
	}

	for _, c := range cases {
		actual := c.input.Rune()
		if actual != c.expected {
			t.Errorf("got %v\nwant %v", actual, c.expected)
		}
	}
}
