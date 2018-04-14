package core

import (
	goast "go/ast"
	"go/token"

	"github.com/droptheplot/yal/yal/ast"
)

func ADD(node ast.Node) goast.Expr {
	e := &goast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.ADD,
	}

	return e
}

func SUB(node ast.Node) goast.Expr {
	return &goast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.SUB,
	}
}

func MUL(node ast.Node) goast.Expr {
	return &goast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.MUL,
	}
}

func QUO(node ast.Node) goast.Expr {
	return &goast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.QUO,
	}
}

func REM(node ast.Node) goast.Expr {
	return &goast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.REM,
	}
}

func EQL(node ast.Node) goast.Expr {
	return &goast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.EQL,
	}
}

func NEQ(node ast.Node) goast.Expr {
	return &goast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.NEQ,
	}
}

func GTR(node ast.Node) goast.Expr {
	return &goast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.GTR,
	}
}

func GEQ(node ast.Node) goast.Expr {
	return &goast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.GEQ,
	}
}

func LSS(node ast.Node) goast.Expr {
	return &goast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.LSS,
	}
}

func LEQ(node ast.Node) goast.Expr {
	return &goast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.LEQ,
	}
}

func LOR(node ast.Node) goast.Expr {
	return &goast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.LOR,
	}
}

func LAND(node ast.Node) goast.Expr {
	return &goast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.LAND,
	}
}
