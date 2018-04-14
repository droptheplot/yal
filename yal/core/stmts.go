package core

import (
	goast "go/ast"
	"go/token"

	"github.com/droptheplot/yal/yal/ast"
)

func IF(node ast.Node) goast.Stmt {
	return &goast.IfStmt{
		Cond: Expr(node.Nodes[0]),
		Body: &goast.BlockStmt{
			List: []goast.Stmt{Stmt(node.Nodes[1])},
		},
	}
}

func VAR(node ast.Node) goast.Stmt {
	return &goast.DeclStmt{
		Decl: &goast.GenDecl{
			Tok: token.VAR,
			Specs: []goast.Spec{
				&goast.ValueSpec{Names: []*goast.Ident{goast.NewIdent(node.Nodes[0].Atom)}, Type: goast.NewIdent(node.Nodes[1].Atom)},
			},
		},
	}
}

func RETURN(node ast.Node) goast.Stmt {
	var exprs []goast.Expr

	for i, _ := range node.Nodes {
		exprs = append(exprs, Expr(node.Nodes[i]))
	}

	return &goast.ReturnStmt{Results: exprs}
}

func ASSIGN(node ast.Node) goast.Stmt {
	var lhs []goast.Expr
	var rhs []goast.Expr

	for i := 0; i < len(node.Nodes); i = i + 2 {
		lhs = append(lhs, Expr(node.Nodes[i]))
		rhs = append(rhs, Expr(node.Nodes[i+1]))
	}

	return &goast.AssignStmt{Tok: token.ASSIGN, Lhs: lhs, Rhs: rhs}
}

func SWITCH(node ast.Node) goast.Stmt {
	var list []goast.Stmt

	for i := 1; i < len(node.Nodes); i = i + 2 {
		var cl []goast.Expr

		if !isDefault(node.Nodes[i]) {
			cl = append(cl, Expr(node.Nodes[i]))
		}

		list = append(list, &goast.CaseClause{
			List: cl,
			Body: []goast.Stmt{
				Stmt(node.Nodes[i+1]),
			},
		})
	}

	return &goast.SwitchStmt{
		Tag:  Expr(node.Nodes[0]),
		Body: &goast.BlockStmt{List: list},
	}
}
