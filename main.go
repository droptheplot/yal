package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"

	"github.com/droptheplot/yal/parser"
	"github.com/droptheplot/yal/tokenizer"
	"github.com/droptheplot/yal/yal"
)

func main() {
	var path string

	flag.StringVar(&path, "path", "", "Path to file.")
	flag.Parse()

	if path == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	src, _ := ioutil.ReadFile(path)

	file := yal.New(&tokenizer.Tokenizer{}, &parser.Parser{}).Run(src)

	PrintFile(file)
}

func PrintFile(file *ast.File) {
	buffer := &bytes.Buffer{}

	printer.Fprint(buffer, token.NewFileSet(), file)

	fmt.Println(buffer)
}
