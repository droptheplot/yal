package core

import (
	goast "go/ast"
	"go/token"

	"github.com/droptheplot/yal/ast"
)

// IF returns "if" statement.
//  (if a b)
// Becomes:
//  if a {
//    b
//  }
func IF(node ast.Node) goast.Stmt {
	return &goast.IfStmt{
		Cond: Expr(node.Nodes[0]),
		Body: &goast.BlockStmt{
			List: []goast.Stmt{Stmt(node.Nodes[1])},
		},
	}
}

// VAR returns "var" statement.
//  (var a b)
// Becomes:
//  var a b
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

// RETURN returns "return" statement.
// Accepts multiple arguments.
//  (return a b c)
// Becomes:
//  return a, b, c
func RETURN(node ast.Node) goast.Stmt {
	var exprs []goast.Expr

	for i, _ := range node.Nodes {
		exprs = append(exprs, Expr(node.Nodes[i]))
	}

	return &goast.ReturnStmt{Results: exprs}
}

// ASSIGN returns "=" statement.
// Accepts multiple arguments.
//  (= a b c d)
// Becomes:
//  a, c = b, d
func ASSIGN(node ast.Node) goast.Stmt {
	var lhs []goast.Expr
	var rhs []goast.Expr

	for i := 0; i < len(node.Nodes); i = i + 2 {
		lhs = append(lhs, Expr(node.Nodes[i]))
		rhs = append(rhs, Expr(node.Nodes[i+1]))
	}

	return &goast.AssignStmt{Tok: token.ASSIGN, Lhs: lhs, Rhs: rhs}
}

// SWITCH returns "switch" statement.
// Accepts multiple arguments.
//  (switch a b c)
// Becomes:
//  switch a {
//    case b: c
//  }
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
