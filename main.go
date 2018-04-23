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
	var debugTokens, debugNodes bool
	var path string

	flag.StringVar(&path, "path", "", "Path to file.")
	flag.BoolVar(&debugTokens, "tokens", false, "Print tokens.")
	flag.BoolVar(&debugNodes, "nodes", false, "Print nodes.")
	flag.Parse()

	if path == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	src, _ := ioutil.ReadFile(path)
	tokens := tokenizer.Tokenize(src)
	node, _ := parser.Parse(tokens)
	file := yal.File(node)

	if debugTokens {
		fmt.Printf("%#v\n\n", tokens)
	}

	if debugNodes {
		fmt.Printf("%# v\n\n", pretty.Formatter(node))
	}

	pckg, _ := ast.NewPackage(token.NewFileSet(), map[string]*ast.File{"main": file}, nil, nil)

	PrintFile(pckg.Files["main"])
}

func PrintFile(file *ast.File) {
	buffer := &bytes.Buffer{}

	printer.Fprint(buffer, token.NewFileSet(), file)

	fmt.Println(buffer)
}
