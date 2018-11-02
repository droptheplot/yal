package main

import (
	"flag"
	"fmt"
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

	src, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file := yal.New(&tokenizer.Tokenizer{}, &parser.Parser{}).Run(src)

	buf, err := yal.Buffer(file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(buf)
}
