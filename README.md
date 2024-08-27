# Morse

Basic Morse code decoder.

## Usage

`morse` reads from stdin, expecting dots as `.`, dashes as `-`, spaces ` ` between letters and tabs or repeat spaces between words.

```sh
$ echo "-- --- .-. ... .\t-.-. --- -.. ." | morse
MORSE CODE
```
