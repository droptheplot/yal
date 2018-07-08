package yal

import (
	"go/ast"
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
	}

	Stmts = map[string]func(Node) ast.Stmt{
		"if":     IF,
		"var":    VAR,
		"return": RETURN,
		"=":      ASSIGN,
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
