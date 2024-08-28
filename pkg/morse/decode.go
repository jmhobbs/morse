package morse

import (
	"bufio"
	"fmt"
	"io"
)

func Decode(in io.Reader, out io.Writer) error {
	var (
		err  error
		c    rune
		char Morse
	)

	r := bufio.NewReader(in)
	for {
		if c, _, err = r.ReadRune(); err != nil {
			fmt.Fprintf(out, "%c", char.Rune())
			if err != io.EOF {
				return err
			}
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
				fmt.Fprintf(out, "%c", char.Rune())
				char = 0
			}
			fmt.Fprint(out, " ")
			break
		case ' ':
			if char == 0 {
				fmt.Fprint(out, " ")
			} else {
				fmt.Fprintf(out, "%c", char.Rune())
				char = 0
			}
			break
		case '\n':
			fmt.Fprintf(out, "%c", char.Rune())
			char = 0
			fmt.Fprintln(out)
			break
		default:
			return fmt.Errorf("Invalid input character: %c\n", c)
		}
	}

	return nil
}
