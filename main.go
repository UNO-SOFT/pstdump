package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/UNO-SOFT/pstdump/parse"
	"github.com/pkg/errors"
)

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}

func Main() error {
	flagFlat := flag.Bool("flat", false, "flatten folder structure")
	flagVerbose := flag.Bool("v", false, "verbose logging")
	flagFn := flag.String("filename", "{{urlquery .Folder}}-{{.ArticleNumber}}.eml", "template of file name")
	flag.Parse()
	destDir := flag.Arg(0)
	if destDir == "" {
		destDir = "."
	}
	if destDir != "." {
		os.MkdirAll(destDir, 0755)
	}
	tmpl := template.Must(template.New("").Parse(*flagFn))

	seen := make(map[string]struct{})
	var buf bytes.Buffer
	r := io.Reader(os.Stdin)
	return parse.Parse(r, func(eml *parse.Email) error {
		dn := destDir
		buf.Reset()
		if err := tmpl.Execute(&buf, eml); err != nil {
			return errors.Wrapf(err, "%+v", eml)
		}
		fn := buf.String()
		if !*flagFlat {
			dn = filepath.Join(dn, eml.Folder)
			if _, ok := seen[dn]; !ok {
				os.MkdirAll(dn, 0755)
				seen[dn] = struct{}{}
			}
			fn = filepath.Join(dn, fn)
		}
		if *flagVerbose {
			log.Println(fn)
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
