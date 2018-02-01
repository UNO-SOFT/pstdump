package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/UNO-SOFT/pstdump/parse"
)

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}

func Main() error {
	flag.Parse()
	r := io.Reader(os.Stdin)
	return parse.Parse(r, func(eml *parse.Email) error {
		eml.WriteTo(os.Stdout)
		return nil
	})
}
