package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jmhobbs/morse/pkg/morse"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "usage: morse [encode|decode]")
	}
	flag.Parse()
	action := flag.Arg(0)

	if action == "encode" {
		err := morse.Encode(os.Stdin, os.Stdout)
		if err != nil {
			log.Fatal(err)
		}
	} else if action == "decode" || action == "" {
		err := morse.Decode(os.Stdin, os.Stdout)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}

}
