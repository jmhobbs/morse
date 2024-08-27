package morse

import "strings"

type Morse uint16

const (
	Dot  Morse = 0b01
	Dash Morse = 0b11
)

func New() *Morse {
	m := Morse(0)
	return &m
}

func (m Morse) ToMorseString() string {
	buf := []string{}
	for {
		if m&Dash == Dash {
			buf = append(buf, "-")
		} else if m&Dot == Dot {
			buf = append(buf, ".")
		} else {
			break
		}
		m >>= 2
	}
	return strings.Join(buf, "")
}

func (m Morse) ToString() string {
	return morseToChar[m]
}

func (m Morse) Dot() Morse {
	return (m << 2) | Dot
}

func (m Morse) Dash() Morse {
	return (m << 2) | Dash
}
