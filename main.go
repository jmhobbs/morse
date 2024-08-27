package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// 2 bytes max, uint16
// 01 - dot
// 11 - dash
var morseToChar = map[uint16]string{
	0b0111:       "A",
	0b11010101:   "B",
	0b11011101:   "C",
	0b110101:     "D",
	0b01:         "E",
	0b01011101:   "F",
	0b111101:     "G",
	0b01010101:   "H",
	0b0101:       "I",
	0b01111111:   "J",
	0b110111:     "K",
	0b01110101:   "L",
	0b1111:       "M",
	0b1101:       "N",
	0b111111:     "O",
	0b01111101:   "P",
	0b11110111:   "Q",
	0b011101:     "R",
	0b010101:     "S",
	0b11:         "T",
	0b010111:     "U",
	0b01010111:   "V",
	0b011111:     "W",
	0b11010111:   "X",
	0b11011111:   "Y",
	0b11110101:   "Z",
	0b1111111111: "0",
	0b0111111111: "1",
	0b0101111111: "2",
	0b0101011111: "3",
	0b0101010111: "4",
	0b0101010101: "5",
	0b1101010101: "6",
	0b1111010101: "7",
	0b1111110101: "8",
	0b1111111101: "9",
}

var charToMorse = map[string]uint16{}

func init() {
	for morse, char := range morseToChar {
		charToMorse[char] = morse
	}
}

func morseToString(morse uint16) string {
	buf := []string{}
	for {
		if morse&0b11 == 0b11 {
			buf = append(buf, "-")
		} else if morse&0b01 == 0b01 {
			buf = append(buf, ".")
		} else {
			break
		}
		morse >>= 2
	}
	return strings.Join(buf, "")
}

func main() {
	var (
		n    int
		err  error
		buf  []byte = make([]byte, 1)
		char uint16
	)

	for {
		n, err = os.Stdin.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			fmt.Print(morseToChar[char])
			break
		}
		if n == 0 {
			break
		}

		switch buf[0] {
		case '-':
			char = (char << 2) | 0b11
			break
		case '.':
			char = (char << 2) | 0b01
			break
		case '\t':
			if char > 0 {
				fmt.Print(morseToChar[char])
				char = 0
			}
			fmt.Printf(" ")
			break
		case ' ':
			if char == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(morseToChar[char])
				char = 0
			}
			break
		case '\n':
			fmt.Print(morseToChar[char])
			char = 0
			fmt.Println()
			break
		default:
			log.Printf("Invalid input character: %c\n", buf[0])
		}
	}
}
