package main

import (
	"fmt"
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
	return parse.Parse(os.Stdin, func(eml *parse.Email) error {
		fmt.Println(eml)
		return nil
	})
}
