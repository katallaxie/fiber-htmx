package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/ettle/strcase"
	"github.com/katallaxie/pkg/logx"
	"github.com/spf13/pflag"
	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
)

type flags struct {
	Output string
}

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	//nolint:errcheck
	logx.RedirectStdLog(logx.LogSink)

	f := &flags{}

	pflag.StringVar(&f.Output, "output", f.Output, "output")
	pflag.Parse()

	// Load template
	tmpl, err := template.ParseFiles("cmd/tailwind/classes.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	// Load the CSS file
	tf, err := os.Open("./cmd/tailwind/tailwind.css")
	if err != nil {
		log.Fatal(err)
	}

	l := css.NewLexer(parse.NewInput(tf))

	vars := make(map[string]string)

OUTTER:
	for {
		tt, text := l.Next()
		//nolint:exhaustive
		switch tt {
		case css.IdentToken:
			name := strings.ReplaceAll(string(text), "\\/", "_")
			name = strings.ReplaceAll(name, "\\", "_")
			name = strings.ReplaceAll(name, "%", "Percent")
			name = strings.ReplaceAll(name, ".", "Fraction")
			name = strcase.ToPascal(name)

			ref := strings.ReplaceAll(string(text), "\\/", "\\\\/")
			ref = strings.ReplaceAll(ref, "\\%", "\\\\%")
			ref = strings.ReplaceAll(ref, "\\.", "\\\\.")

			vars[name] = ref
		case css.ErrorToken:
			break OUTTER
		}
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, vars)
	if err != nil {
		log.Fatal(err)
	}

	out, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	file, _ := os.Create("components/tailwind/classes.gen.go")
	defer file.Close()

	_, err = file.Write(out)
	if err != nil {
		log.Print(err)
	}
}
