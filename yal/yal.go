package yal

import (
	"go/ast"
)

var Exprs map[string]func(Node) ast.Expr
var Stmts map[string]func(Node) ast.Stmt

func init() {
	Exprs = map[string]func(Node) ast.Expr{
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

	Stmts = map[string]func(Node) ast.Stmt{
		"if":     IF,
		"var":    VAR,
		"return": RETURN,
		"=":      ASSIGN,
		"switch": SWITCH,
	}
}

func isExpr(node Node) bool {
	_, ok := Exprs[node.Atom]

	return ok
}

func isStmt(node Node) bool {
	_, ok := Stmts[node.Atom]

	return ok
}

func isFunc(node Node) bool {
	return node.Atom == "func"
}

func isPackage(node Node) bool {
	return node.Atom == "package"
}

func isDefault(node Node) bool {
	return node.Atom == "default"
}

func File(node Node) *ast.File {
	var name *ast.Ident
	var decls []ast.Decl

	for i := range node.Nodes {
		if isFunc(node.Nodes[i]) {
			decls = append(decls, Func(node.Nodes[i]))
		} else if isPackage(node.Nodes[i]) {
			name = ast.NewIdent(node.Nodes[i].Nodes[0].Atom)
		}
	}

	return &ast.File{
		Name:  name,
		Decls: decls,
		Scope: &ast.Scope{},
	}
}

func Func(node Node) ast.Decl {
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

func FieldList(node Node) *ast.FieldList {
	var list []*ast.Field

	if node.Atom == "" {
		for i := range node.Nodes {
			list = append(list, Field(node.Nodes[i]))
		}
	} else {
		list = append(list, Field(node))
	}

	return &ast.FieldList{List: list}
}

func Field(node Node) *ast.Field {
	if len(node.Nodes) == 0 {
		return &ast.Field{Type: ast.NewIdent(node.Atom)}
	}

	return &ast.Field{Names: []*ast.Ident{ast.NewIdent(node.Atom)}, Type: ast.NewIdent(node.Nodes[0].Atom)}
}

func Stmt(node Node) ast.Stmt {
	if isStmt(node) {
		return Stmts[node.Atom](node)
	}

	return &ast.ExprStmt{X: Expr(node)}
}

func Expr(node Node) ast.Expr {
	if len(node.Nodes) == 0 {
		return &ast.BasicLit{Value: node.Atom}
	} else if isExpr(node) {
		return Exprs[node.Atom](node)
	} else {
		var args []ast.Expr

		for i := range node.Nodes {
			args = append(args, Expr(node.Nodes[i]))
		}

		e := &ast.CallExpr{
			Fun:  ast.NewIdent(node.Atom),
			Args: args,
		}

		return e
	}
}
