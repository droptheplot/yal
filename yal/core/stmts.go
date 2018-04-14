package core

import (
	"go/ast"
	"go/token"

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

func VAR(node yal.Node) ast.Stmt {
	return &ast.DeclStmt{
		Decl: &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{Names: []*ast.Ident{ast.NewIdent(node.Nodes[0].Atom)}, Type: ast.NewIdent(node.Nodes[1].Atom)},
			},
		},
	}
}

func RETURN(node yal.Node) ast.Stmt {
	var exprs []ast.Expr

	for i, _ := range node.Nodes {
		exprs = append(exprs, Expr(node.Nodes[i]))
	}

	return &ast.ReturnStmt{Results: exprs}
}

func ASSIGN(node yal.Node) ast.Stmt {
	var lhs []ast.Expr
	var rhs []ast.Expr

	for i := 0; i < len(node.Nodes); i = i + 2 {
		lhs = append(lhs, Expr(node.Nodes[i]))
		rhs = append(rhs, Expr(node.Nodes[i+1]))
	}

	return &ast.AssignStmt{Tok: token.ASSIGN, Lhs: lhs, Rhs: rhs}
}

func SWITCH(node yal.Node) ast.Stmt {
	var list []ast.Stmt

	for i := 1; i < len(node.Nodes); i = i + 2 {
		var cl []ast.Expr

		if !isDefault(node.Nodes[i]) {
			cl = append(cl, Expr(node.Nodes[i]))
		}

		list = append(list, &ast.CaseClause{
			List: cl,
			Body: []ast.Stmt{
				Stmt(node.Nodes[i+1]),
			},
		})
	}

	return &ast.SwitchStmt{
		Tag:  Expr(node.Nodes[0]),
		Body: &ast.BlockStmt{List: list},
	}
}
