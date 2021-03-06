package yal

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
)

type Tokenizer interface {
	Tokenize(src []byte) []string
}

type Parser interface {
	Parse(tokens []string) Node
}

type Yal struct {
	tokenizer Tokenizer
	parser    Parser
}

var Exprs map[string]func(Node) ast.Expr
var Stmts map[string]func(Node) ast.Stmt

func init() {
	Exprs = map[string]func(Node) ast.Expr{
		"+":     ADD,
		"-":     SUB,
		"*":     MUL,
		"/":     QUO,
		"%":     REM,
		"==":    EQL,
		"!=":    NEQ,
		">":     GTR,
		">=":    GEQ,
		"<":     LSS,
		"<=":    LEQ,
		"||":    LOR,
		"&&":    LAND,
		"slice": SLICE,
		"array": ARRAY,
		"map":   MAP,
	}

	Stmts = map[string]func(Node) ast.Stmt{
		"if":     IF,
		"var":    VAR,
		"return": RETURN,
		"=":      ASSIGN,
		":=":     DEFINE,
		"switch": SWITCH,
	}
}

func New(t Tokenizer, p Parser) *Yal {
	return &Yal{
		tokenizer: t,
		parser:    p,
	}
}

func (y *Yal) Run(src []byte) *ast.File {
	tokens := y.tokenizer.Tokenize(src)
	node := y.parser.Parse(tokens)
	file := node.File()

	return file
}

func Buffer(file *ast.File) (*bytes.Buffer, error) {
	buffer := &bytes.Buffer{}

	err := printer.Fprint(buffer, token.NewFileSet(), file)

	if err != nil {
		return nil, err
	}

	return buffer, nil
}
