package core

import (
	"go/ast"

	"github.com/droptheplot/yal/yal"
)

var Exprs map[string]func(yal.Node) ast.Expr
var Stmts map[string]func(yal.Node) ast.Stmt

func init() {
	Exprs = map[string]func(yal.Node) ast.Expr{
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

	Stmts = map[string]func(yal.Node) ast.Stmt{
		"if":     IF,
		"var":    VAR,
		"return": RETURN,
		"=":      ASSIGN,
	}
}

func isExpr(node yal.Node) bool {
	_, ok := Exprs[node.Atom]

	return ok
}

func isStmt(node yal.Node) bool {
	_, ok := Stmts[node.Atom]

	return ok
}

func isFunc(node yal.Node) bool {
	return node.Atom == "func"
}

func isPackage(node yal.Node) bool {
	return node.Atom == "package"
}

func File(node yal.Node) []ast.Decl {
	var result []ast.Decl

	for _, n := range node.Nodes {
		if isFunc(n) {
			result = append(result, Func(n))
		}
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
	if isStmt(node) {
		return Stmts[node.Atom](node)
	} else {
		return &ast.ExprStmt{X: Expr(node)}
	}
}

func Expr(node yal.Node) ast.Expr {
	if len(node.Nodes) == 0 {
		return &ast.BasicLit{Value: node.Atom}
	} else if isExpr(node) {
		return Exprs[node.Atom](node)
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

func Name(node yal.Node) *ast.Ident {
	var name string

	for _, n := range node.Nodes {
		if isPackage(n) {
			name = n.Nodes[0].Atom
		}
	}

	return ast.NewIdent(name)
}
