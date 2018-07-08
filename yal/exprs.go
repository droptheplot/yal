package yal

import (
	"go/ast"
	"go/token"
)

// ADD returns expression for addition.
//  (+ a b)
func ADD(node Node) ast.Expr {
	e := &ast.BinaryExpr{
		X:  node.Nodes[0].Expr(),
		Y:  node.Nodes[1].Expr(),
		Op: token.ADD,
	}

	return e
}

// SUB returns expression for subtraction.
//  (- a b)
func SUB(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  node.Nodes[0].Expr(),
		Y:  node.Nodes[1].Expr(),
		Op: token.SUB,
	}
}

// MUL returns expression for multiplication.
//  (* a b)
func MUL(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  node.Nodes[0].Expr(),
		Y:  node.Nodes[1].Expr(),
		Op: token.MUL,
	}
}

// QUO returns expression for division.
//  (/ a b)
func QUO(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  node.Nodes[0].Expr(),
		Y:  node.Nodes[1].Expr(),
		Op: token.QUO,
	}
}

// REM returns expression for remainder.
//  (% a b)
func REM(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  node.Nodes[0].Expr(),
		Y:  node.Nodes[1].Expr(),
		Op: token.REM,
	}
}

// EQL returns expression for equality.
//  (== a b)
func EQL(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  node.Nodes[0].Expr(),
		Y:  node.Nodes[1].Expr(),
		Op: token.EQL,
	}
}

// NEQ returns expression for inequality.
//  (!= a b)
func NEQ(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  node.Nodes[0].Expr(),
		Y:  node.Nodes[1].Expr(),
		Op: token.NEQ,
	}
}

// GTR returns "greater than" expression.
//  (> a b)
func GTR(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  node.Nodes[0].Expr(),
		Y:  node.Nodes[1].Expr(),
		Op: token.GTR,
	}
}

// GEQ returns "greater than or equal" expression.
//  (>= a b)
func GEQ(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  node.Nodes[0].Expr(),
		Y:  node.Nodes[1].Expr(),
		Op: token.GEQ,
	}
}

// LSS returns "less than" expression.
//  (< a b)
func LSS(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  node.Nodes[0].Expr(),
		Y:  node.Nodes[1].Expr(),
		Op: token.LSS,
	}
}

// LEQ returns "less than or equal" expression.
//  (<= a b)
func LEQ(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  node.Nodes[0].Expr(),
		Y:  node.Nodes[1].Expr(),
		Op: token.LEQ,
	}
}

// LOR returns "or" expression.
//  (<= a b)
func LOR(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  node.Nodes[0].Expr(),
		Y:  node.Nodes[1].Expr(),
		Op: token.LOR,
	}
}

// LAND returns "and" expression.
//  (&& a b)
func LAND(node Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  node.Nodes[0].Expr(),
		Y:  node.Nodes[1].Expr(),
		Op: token.LAND,
	}
}

// SLICE returns composite literal expression for slice.
// First node represents element type.
//  (slice type a b)
func SLICE(node Node) ast.Expr {
	var elts []ast.Expr

	for _, n := range node.Nodes[1:] {
		elts = append(elts, n.Expr())
	}

	return &ast.CompositeLit{
		Type: &ast.ArrayType{
			Elt: ast.NewIdent(node.Nodes[0].Atom),
		},
		Elts: elts,
	}
}

// ARRAY returns composite literal expression for array.
// First node represents element type.
// Second node represents array size.
//  (array type n a b)
func ARRAY(node Node) ast.Expr {
	var elts []ast.Expr

	for _, n := range node.Nodes[1:] {
		elts = append(elts, n.Expr())
	}

	return &ast.CompositeLit{
		Type: &ast.ArrayType{
			Len: node.Nodes[1].Expr(),
			Elt: ast.NewIdent(node.Nodes[0].Atom),
		},
		Elts: elts,
	}
}
