package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jmhobbs/morse/pkg/morse"
)

func main() {
	var (
		n    int
		err  error
		buf  []byte = make([]byte, 1)
		char morse.Morse
	)

	for {
		n, err = os.Stdin.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			fmt.Print(char.ToString())
			break
		}
		if n == 0 {
			break
		}

		switch buf[0] {
		case '-':
			char = char.Dash()
			break
		case '.':
			char = char.Dot()
			break
		case '\t':
			if char > 0 {
				fmt.Print(char.ToString())
				char = 0
			}
			fmt.Printf(" ")
			break
		case ' ':
			if char == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(char.ToString())
				char = 0
			}
			break
		case '\n':
			fmt.Print(char.ToString())
			char = 0
			fmt.Println()
			break
		default:
			log.Printf("Invalid input character: %c\n", buf[0])
		}
	}
}
