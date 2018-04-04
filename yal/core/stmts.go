package core

import (
	"go/ast"

	"github.com/droptheplot/yal/yal"
)

func IF(node yal.Node) ast.Stmt {
	return &ast.IfStmt{
		Cond: Expr(node.Nodes[0]),
		Body: &ast.BlockStmt{
			List: []ast.Stmt{Stmt(node.Nodes[1])},
		},
	}
}
