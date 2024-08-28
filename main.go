package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jmhobbs/morse/pkg/morse"
)

func main() {
	var (
		err  error
		c    rune
		char morse.Morse
	)

	r := bufio.NewReader(os.Stdin)
	for {
		if c, _, err = r.ReadRune(); err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			fmt.Printf("%c", char.Rune())
			break
		}

		switch c {
		case '-':
			char = char.Dash()
			break
		case '.':
			char = char.Dot()
			break
		case '\t':
			if char > 0 {
				fmt.Printf("%c", char.Rune())
				char = 0
			}
			fmt.Printf(" ")
			break
		case ' ':
			if char == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c", char.Rune())
				char = 0
			}
			break
		case '\n':
			fmt.Printf("%c", char.Rune())
			char = 0
			fmt.Println()
			break
		default:
			log.Printf("Invalid input character: %c\n", r)
		}
	}
}
