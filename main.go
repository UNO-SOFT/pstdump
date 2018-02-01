package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/UNO-SOFT/pstdump/parse"
)

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}

func Main() error {
	flagFlat := flag.Bool("flat", false, "flatten folder structure")
	flag.Parse()
	destDir := flag.Arg(0)
	if destDir == "" {
		destDir = "."
	}
	os.MkdirAll(destDir, 0755)

	r := io.Reader(os.Stdin)
	return parse.Parse(r, func(eml *parse.Email) error {
		dn := destDir
		var fn string
		if *flagFlat {
			fn = filepath.Join(dn, fmt.Sprintf("%s-%d.eml", eml.Folder, eml.ArticleNumber))
		} else {
			dn = filepath.Join(dn, eml.Folder)
			os.MkdirAll(dn, 0755)
			fn = filepath.Join(dn, fmt.Sprintf("%d.eml", eml.ArticleNumber))
		}
		fh, err := os.Create(fn)
		if err != nil {
			return err
		}
		_, err = eml.WriteTo(fh)
		if closeErr := fh.Close(); err == nil {
			return closeErr
		}
		return err
	})
}
