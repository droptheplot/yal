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

	fset := token.NewFileSet()
	f := yal.New(src)

	if debug {
		fmt.Printf("%# v\n\n", pretty.Formatter(f))
	}

	pckg, _ := ast.NewPackage(fset, map[string]*ast.File{"main": f}, nil, nil)

	out, _ := GenerateFile(fset, pckg.Files["main"])
	fmt.Println(string(out))
}

func GenerateFile(fset *token.FileSet, file *ast.File) ([]byte, error) {
	var output []byte

	buffer := bytes.NewBuffer(output)

	if err := printer.Fprint(buffer, fset, file); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
