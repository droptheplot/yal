package main

import (
	"go/ast"
	"go/token"
)

var coreFuns map[string]interface{}

func init() {
	coreFuns = map[string]interface{}{
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

func isCoreFun(atom string) bool {
	_, ok := coreFuns[atom]

	return ok
}

func ADD(node Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  evalExpr(node.Nodes[0]),
		Y:  evalExpr(node.Nodes[1]),
		Op: token.ADD,
	}

	return e
}

func SUB(node Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  evalExpr(node.Nodes[0]),
		Y:  evalExpr(node.Nodes[1]),
		Op: token.SUB,
	}

	return e
}

func MUL(node Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  evalExpr(node.Nodes[0]),
		Y:  evalExpr(node.Nodes[1]),
		Op: token.MUL,
	}

	return e
}

func QUO(node Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  evalExpr(node.Nodes[0]),
		Y:  evalExpr(node.Nodes[1]),
		Op: token.QUO,
	}

	return e
}
func REM(node Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  evalExpr(node.Nodes[0]),
		Y:  evalExpr(node.Nodes[1]),
		Op: token.REM,
	}

	return e
}
func EQL(node Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  evalExpr(node.Nodes[0]),
		Y:  evalExpr(node.Nodes[1]),
		Op: token.EQL,
	}

	return e
}

func GTR(node Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  evalExpr(node.Nodes[0]),
		Y:  evalExpr(node.Nodes[1]),
		Op: token.GTR,
	}

	return e
}

func GEQ(node Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  evalExpr(node.Nodes[0]),
		Y:  evalExpr(node.Nodes[1]),
		Op: token.GEQ,
	}

	return e
}

func LSS(node Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  evalExpr(node.Nodes[0]),
		Y:  evalExpr(node.Nodes[1]),
		Op: token.LSS,
	}

	return e
}

func LEQ(node Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  evalExpr(node.Nodes[0]),
		Y:  evalExpr(node.Nodes[1]),
		Op: token.LEQ,
	}

	return e
}
