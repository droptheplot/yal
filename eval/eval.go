package eval

import (
	"go/ast"

	"github.com/droptheplot/yal/yal"
)

func File(node yal.Node) []ast.Decl {
	var result []ast.Decl

	for _, n := range node.Nodes {
		result = append(result, Func(n))
	}

	return result
}

func Func(node yal.Node) ast.Decl {
	var bodyList []ast.Stmt

	for _, n := range node.Nodes[3:] {
		bodyList = append(bodyList, Stmt(n))
	}

	result := &ast.FuncDecl{
		Name: ast.NewIdent(node.Nodes[0].Atom),
		Type: &ast.FuncType{
			Params:  FieldList(node.Nodes[1]),
			Results: FieldList(node.Nodes[2]),
		},
		Body: &ast.BlockStmt{
			List: bodyList,
		},
	}

	return result
}

func FieldList(node yal.Node) *ast.FieldList {
	var list []*ast.Field

	if node.Atom == "" {
		for _, n := range node.Nodes {
			list = append(list, Field(n))
		}
	} else {
		list = append(list, Field(node))
	}

	return &ast.FieldList{List: list}
}

func Field(node yal.Node) *ast.Field {
	if len(node.Nodes) == 0 {
		return &ast.Field{Type: ast.NewIdent(node.Atom)}
	}

	return &ast.Field{Names: []*ast.Ident{ast.NewIdent(node.Atom)}, Type: ast.NewIdent(node.Nodes[0].Atom)}
}

func Stmt(node yal.Node) ast.Stmt {
	if isCoreStmt(node) {
		return coreStmts[node.Atom](node)
	} else {
		return &ast.ExprStmt{X: Expr(node)}
	}
}

func Expr(node yal.Node) ast.Expr {
	if len(node.Nodes) == 0 {
		return &ast.BasicLit{Value: node.Atom}
	} else if isCoreExpr(node) {
		return coreExprs[node.Atom](node)
	} else {
		var args []ast.Expr

		for _, n := range node.Nodes {
			args = append(args, Expr(n))
		}

		e := &ast.CallExpr{
			Fun:  ast.NewIdent(node.Atom),
			Args: args,
		}

		return e
	}
}
