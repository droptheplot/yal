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

	"github.com/kr/pretty"
)

func main() {
	var debug bool
	var file string

	flag.StringVar(&file, "file", "", "Path to file")
	flag.BoolVar(&debug, "debug", false, "Debug mode")

	flag.Parse()

	if file == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	src, _ := ioutil.ReadFile(file)

	f := GenerateFile(src)

	if debug {
		fmt.Printf("%# v\n\n", pretty.Formatter(f))
	}

	pckg, _ := ast.NewPackage(token.NewFileSet(), map[string]*ast.File{"main": f}, nil, nil)

	out, _ := PrintFile(pckg.Files["main"])
	fmt.Println(string(out))
}

func GenerateFile(src []byte) *ast.File {
	tokens := tokenizer.Tokenize(src)
	node, _ := parser.Parse(tokens)

	return yal.File(node)
}

func PrintFile(file *ast.File) ([]byte, error) {
	var output []byte

	buffer := bytes.NewBuffer(output)

	if err := printer.Fprint(buffer, token.NewFileSet(), file); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
