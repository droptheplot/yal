package yal

import (
	"go/ast"
	"go/token"
)

// ADD returns expression for addition.
//  (+ a b)
func ADD(node Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.ADD,
	}

	return e
}

// SUB returns expression for subtraction.
//  (- a b)
func SUB(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.SUB,
	}
}

// MUL returns expression for multiplication.
//  (* a b)
func MUL(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.MUL,
	}
}

// QUO returns expression for division.
//  (/ a b)
func QUO(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.QUO,
	}
}

// REM returns expression for remainder.
//  (% a b)
func REM(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.REM,
	}
}

// EQL returns expression for equality.
//  (== a b)
func EQL(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.EQL,
	}
}

// NEQ returns expression for inequality.
//  (!= a b)
func NEQ(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.NEQ,
	}
}

// GTR returns "greater than" expression.
//  (> a b)
func GTR(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.GTR,
	}
}

// GEQ returns "greater than or equal" expression.
//  (>= a b)
func GEQ(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.GEQ,
	}
}

// LSS returns "less than" expression.
//  (< a b)
func LSS(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.LSS,
	}
}

// LEQ returns "less than or equal" expression.
//  (<= a b)
func LEQ(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.LEQ,
	}
}

// LOR returns "or" expression.
//  (<= a b)
func LOR(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.LOR,
	}
}

// LAND returns "and" expression.
//  (&& a b)
func LAND(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.LAND,
	}
}
