package morse

import (
	"bufio"
	"fmt"
	"io"
)

func Encode(in io.Reader, out io.Writer) error {
	var (
		err  error
		c    rune
		char *Morse
	)

	r := bufio.NewReader(in)
	for {
		if c, _, err = r.ReadRune(); err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		if c == ' ' || c == '\n' || c == '\t' {
			fmt.Fprintf(out, "%c", c)
			continue
		}

		char = NewFromRune(c)
		if char == nil {
			return fmt.Errorf("Invalid input character: %c\n", c)
		}

		fmt.Fprintf(out, "%s ", char.String())
	}

	return nil
}
