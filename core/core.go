package core

import (
	goast "go/ast"

	"github.com/droptheplot/yal/ast"
)

var Exprs map[string]func(ast.Node) goast.Expr
var Stmts map[string]func(ast.Node) goast.Stmt

func init() {
	Exprs = map[string]func(ast.Node) goast.Expr{
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

	Stmts = map[string]func(ast.Node) goast.Stmt{
		"if":     IF,
		"var":    VAR,
		"return": RETURN,
		"=":      ASSIGN,
		"switch": SWITCH,
	}
}

func isExpr(node ast.Node) bool {
	_, ok := Exprs[node.Atom]

	return ok
}

func isStmt(node ast.Node) bool {
	_, ok := Stmts[node.Atom]

	return ok
}

func isFunc(node ast.Node) bool {
	return node.Atom == "func"
}

func isPackage(node ast.Node) bool {
	return node.Atom == "package"
}

func isDefault(node ast.Node) bool {
	return node.Atom == "default"
}

func File(node ast.Node) *goast.File {
	var name *goast.Ident
	var decls []goast.Decl

	for i := range node.Nodes {
		if isFunc(node.Nodes[i]) {
			decls = append(decls, Func(node.Nodes[i]))
		} else if isPackage(node.Nodes[i]) {
			name = goast.NewIdent(node.Nodes[i].Nodes[0].Atom)
		}
	}

	return &goast.File{
		Name:  name,
		Decls: decls,
		Scope: &goast.Scope{},
	}
}

func Func(node ast.Node) goast.Decl {
	var bodyList []goast.Stmt

	for _, n := range node.Nodes[3:] {
		bodyList = append(bodyList, Stmt(n))
	}

	result := &goast.FuncDecl{
		Name: goast.NewIdent(node.Nodes[0].Atom),
		Type: &goast.FuncType{
			Params:  FieldList(node.Nodes[1]),
			Results: FieldList(node.Nodes[2]),
		},
		Body: &goast.BlockStmt{
			List: bodyList,
		},
	}

	return result
}

func FieldList(node ast.Node) *goast.FieldList {
	var list []*goast.Field

	if node.Atom == "" {
		for i := range node.Nodes {
			list = append(list, Field(node.Nodes[i]))
		}
	} else {
		list = append(list, Field(node))
	}

	return &goast.FieldList{List: list}
}

func Field(node ast.Node) *goast.Field {
	if len(node.Nodes) == 0 {
		return &goast.Field{Type: goast.NewIdent(node.Atom)}
	}

	return &goast.Field{Names: []*goast.Ident{goast.NewIdent(node.Atom)}, Type: goast.NewIdent(node.Nodes[0].Atom)}
}

func Stmt(node ast.Node) goast.Stmt {
	if isStmt(node) {
		return Stmts[node.Atom](node)
	}

	return &goast.ExprStmt{X: Expr(node)}
}

func Expr(node ast.Node) goast.Expr {
	if len(node.Nodes) == 0 {
		return &goast.BasicLit{Value: node.Atom}
	} else if isExpr(node) {
		return Exprs[node.Atom](node)
	} else {
		var args []goast.Expr

		for i := range node.Nodes {
			args = append(args, Expr(node.Nodes[i]))
		}

		e := &goast.CallExpr{
			Fun:  goast.NewIdent(node.Atom),
			Args: args,
		}

		return e
	}
}
