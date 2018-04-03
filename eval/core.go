package eval

import (
	"go/ast"
	"go/token"

	"github.com/droptheplot/yal/yal"
)

var coreFuncs map[string]interface{}

func init() {
	coreFuncs = map[string]interface{}{
		"+":  ADD,
		"-":  SUB,
		"*":  MUL,
		"/":  QUO,
		"%":  REM,
		"==": EQL,
		">":  GTR,
		">=": GEQ,
		"<":  LSS,
		"<=": LEQ,
	}
}

func isCoreFunc(atom string) bool {
	_, ok := coreFuncs[atom]

	return ok
}

func ADD(node yal.Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.ADD,
	}

	return e
}

func SUB(node yal.Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.SUB,
	}

	return e
}

func MUL(node yal.Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.MUL,
	}

	return e
}

func QUO(node yal.Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.QUO,
	}

	return e
}
func REM(node yal.Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.REM,
	}

	return e
}
func EQL(node yal.Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.EQL,
	}

	return e
}

func GTR(node yal.Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.GTR,
	}

	return e
}

func GEQ(node yal.Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.GEQ,
	}

	return e
}

func LSS(node yal.Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.LSS,
	}

	return e
}

func LEQ(node yal.Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.LEQ,
	}

	return e
}
