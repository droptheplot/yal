package eval

import (
	"go/ast"
	"go/token"

	"github.com/droptheplot/yal/yal"
)

var coreExprs map[string]func(yal.Node) ast.Expr
var coreStmts map[string]func(yal.Node) ast.Stmt

func init() {
	coreExprs = map[string]func(yal.Node) ast.Expr{
		"+":  ADD,
		"-":  SUB,
		"*":  MUL,
		"/":  QUO,
		"%":  REM,
		"==": EQL,
		"!=": NEQ,
		">":  GTR,
		">=": GEQ,
		"<":  LSS,
		"<=": LEQ,
		"||": LOR,
		"&&": LAND,
	}

	coreStmts = map[string]func(yal.Node) ast.Stmt{
		"if": IF,
	}
}

func isCoreExpr(node yal.Node) bool {
	_, ok := coreExprs[node.Atom]

	return ok
}

func isCoreStmt(node yal.Node) bool {
	_, ok := coreStmts[node.Atom]

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
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.SUB,
	}
}

func MUL(node yal.Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.MUL,
	}
}

func QUO(node yal.Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.QUO,
	}
}

func REM(node yal.Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.REM,
	}
}

func EQL(node yal.Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.EQL,
	}
}

func NEQ(node yal.Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.NEQ,
	}
}

func GTR(node yal.Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.GTR,
	}
}

func GEQ(node yal.Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.GEQ,
	}
}

func LSS(node yal.Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.LSS,
	}
}

func LEQ(node yal.Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.LEQ,
	}
}

func LOR(node yal.Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.LOR,
	}
}

func LAND(node yal.Node) ast.Expr {
	return &ast.BinaryExpr{
		X:  Expr(node.Nodes[0]),
		Y:  Expr(node.Nodes[1]),
		Op: token.LAND,
	}
}

func IF(node yal.Node) ast.Stmt {
	return &ast.IfStmt{
		Cond: Expr(node.Nodes[0]),
		Body: &ast.BlockStmt{
			List: []ast.Stmt{Stmt(node.Nodes[1])},
		},
	}
}
