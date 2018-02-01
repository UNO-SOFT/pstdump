package main

import (
	"flag"
	"fmt"
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
	log.Printf("Start parsing %v", r)
	return parse.Parse(r, func(eml *parse.Email) error {
		fmt.Println("===========================================")
		eml.WriteTo(os.Stdout)
		return nil
	})
}
