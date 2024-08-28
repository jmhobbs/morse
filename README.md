# Morse

Basic Morse code decoder.

## Usage

`morse` reads from stdin, expecting dots as `.`, dashes as `-`, spaces ` ` between letters and tabs or repeat spaces between words.

```
usage: morse [encode|decode]
```

If neither `encode` nor `decode` is provided, `morse` will default to decoding.

```sh
$ echo "-- --- .-. ... .\t-.-. --- -.. ." | morse
MORSE CODE
```

```sh
$ cho "HELLO WORLD" | morse encode
.... . ..-. ..-. ---  --. --- .-. ..-. ..-
```
