package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var alphabet = map[string]string{
	".-":    "A",
	"-...":  "B",
	"-.-.":  "C",
	"-..":   "D",
	".":     "E",
	"..-.":  "F",
	"--.":   "G",
	"....":  "H",
	"..":    "I",
	".---":  "J",
	"-.-":   "K",
	".-..":  "L",
	"--":    "M",
	"-.":    "N",
	"---":   "O",
	".--.":  "P",
	"--.-":  "Q",
	".-.":   "R",
	"...":   "S",
	"-":     "T",
	"..-":   "U",
	"...-":  "V",
	".--":   "W",
	"-..-":  "X",
	"-.--":  "Y",
	"--..":  "Z",
	"-----": "0",
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",
}

func main() {
	var (
		n    int
		err  error
		buf  []byte = make([]byte, 1)
		char []byte
	)

	for {
		n, err = os.Stdin.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			fmt.Print(alphabet[string(char)])
			break
		}
		if n == 0 {
			break
		}

		switch buf[0] {
		case '-':
			char = append(char, '-')
			break
		case '.':
			char = append(char, '.')
			break
		case '\t':
			if len(char) > 0 {
				fmt.Print(alphabet[string(char)])
				char = nil
			}
			fmt.Printf(" ")
			break
		case ' ':
			if len(char) == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(alphabet[string(char)])
				char = nil
			}
			break
		case '\n':
			fmt.Print(alphabet[string(char)])
			char = nil
			fmt.Println()
			break
		default:
			log.Printf("Invalid character: %c\n", buf[0])
		}
	}
}
