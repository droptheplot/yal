package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kr/pretty"
)

type Node struct {
	Atom  string
	Nodes []Node
}

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

	tokens := tokenize(string(src))
	yalAST, _ := parse(tokens)

	if debug {
		fmt.Printf("%# v\n\n", pretty.Formatter(yalAST))
	}

	fset := token.NewFileSet()

	mainFile := &ast.File{
		Name:  ast.NewIdent("main"),
		Decls: evalFile(yalAST),
		Scope: &ast.Scope{},
	}

	pckg, _ := ast.NewPackage(fset, map[string]*ast.File{"main": mainFile}, nil, nil)

	out, _ := GenerateFile(fset, pckg.Files["main"])
	fmt.Println(string(out))
}

func tokenize(src string) []string {
	var tokens []string
	var char rune
	var err error = nil
	var inString bool = false
	var atom []rune

	reader := bufio.NewReader(strings.NewReader(src))

	for err == nil {
		char, _, err = reader.ReadRune()

		if inString {
			if isQuote(char) {
				inString = false
				atom = append(atom, char)
				tokens = append(tokens, string(atom))
				atom = atom[:0]
			} else {
				atom = append(atom, char)
			}

			continue
		}

		switch {
		case isLParen(char):
			tokens = append(tokens, string(char))
		case isRParen(char):
			applyAtom(&tokens, &atom)
			tokens = append(tokens, string(char))
		case isLiteral(char):
			atom = append(atom, char)
		case isQuote(char):
			inString = true
			atom = append(atom, char)
		case isWhitespace(char):
			applyAtom(&tokens, &atom)
		}
	}

	return tokens
}

func isLParen(r rune) bool {
	return r == '('
}

func isRParen(r rune) bool {
	return r == ')'
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\n'
}

func isQuote(r rune) bool {
	return r == '"'
}

func isLiteral(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') ||
		r == '%' || r == '+' || r == '-' || r == '/' || r == '*' ||
		r == '=' || r == '>' || r == '<' || r == '.'
}

func applyAtom(tokens *[]string, atom *[]rune) {
	if len(*atom) > 0 {
		*tokens = append(*tokens, string(*atom))
		*atom = (*atom)[:0]
	}
}

func parse(tokens []string) (Node, int) {
	var token string
	var size int = len(tokens)
	var i int = 0

	node := Node{}

	for i < size {
		token = tokens[i]

		if token == "(" {
			tmp, move := parse(tokens[i+1:])
			node.Nodes = append(node.Nodes, tmp)
			i += move + 1
		} else if token == ")" {
			return node, i
		} else if node.Atom == "" {
			node.Atom = token
		} else {
			var valueNode = Node{Atom: token}
			node.Nodes = append(node.Nodes, valueNode)
		}

		i++
	}

	return node, size
}

func evalFile(node Node) []ast.Decl {
	var result []ast.Decl

	for _, n := range node.Nodes {
		result = append(result, evalFunc(n))
	}

	return result
}

func evalFunc(node Node) ast.Decl {
	var bodyList []ast.Stmt

	for _, n := range node.Nodes[3:] {
		bodyList = append(bodyList, evalStmt(n))
	}

	result := &ast.FuncDecl{
		Name: ast.NewIdent(node.Nodes[0].Atom),
		Type: &ast.FuncType{
			Params:  evalFieldList(node.Nodes[1]),
			Results: evalFieldList(node.Nodes[2]),
		},
		Body: &ast.BlockStmt{
			List: bodyList,
		},
	}

	return result
}

func evalFieldList(node Node) *ast.FieldList {
	var list []*ast.Field

	if node.Atom == "" {
		for _, n := range node.Nodes {
			list = append(list, evalField(n))
		}
	} else {
		list = append(list, evalField(node))
	}

	return &ast.FieldList{List: list}
}

func evalField(node Node) *ast.Field {
	if len(node.Nodes) == 0 {
		return &ast.Field{Type: ast.NewIdent(node.Atom)}
	}

	return &ast.Field{Names: []*ast.Ident{ast.NewIdent(node.Atom)}, Type: ast.NewIdent(node.Nodes[0].Atom)}
}

func evalStmt(node Node) ast.Stmt {
	return &ast.ExprStmt{X: evalExpr(node)}
}

func evalExpr(node Node) ast.Expr {
	if len(node.Nodes) == 0 {
		return &ast.BasicLit{Value: node.Atom}
	} else if isCoreFun(node.Atom) {
		return coreFuns[node.Atom].(func(node Node) ast.Expr)(node)
	} else {
		var args []ast.Expr

		for _, n := range node.Nodes {
			args = append(args, evalExpr(n))
		}

		e := &ast.CallExpr{
			Fun:  ast.NewIdent(node.Atom),
			Args: args,
		}

		return e
	}
}

func GenerateFile(fset *token.FileSet, file *ast.File) ([]byte, error) {
	var output []byte

	buffer := bytes.NewBuffer(output)

	if err := printer.Fprint(buffer, fset, file); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
