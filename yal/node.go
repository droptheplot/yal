package yal

import (
	"go/ast"
	"go/token"
)

type Node struct {
	Atom  string
	Nodes []Node
}

func (node Node) isExpr() bool {
	_, ok := Exprs[node.Atom]

	return ok
}

func (node Node) isStmt() bool {
	_, ok := Stmts[node.Atom]

	return ok
}

func (node Node) isFunc() bool {
	return node.Atom == "func"
}

func (node Node) isPackage() bool {
	return node.Atom == "package"
}

func (node Node) isDefault() bool {
	return node.Atom == "default"
}

func (node Node) isImport() bool {
	return node.Atom == "import"
}

func (node Node) File() *ast.File {
	var name *ast.Ident
	var decls []ast.Decl
	var specs []ast.Spec

	imports := &ast.GenDecl{Tok: token.IMPORT, Specs: specs}

	decls = append(decls, imports)

	for i := range node.Nodes {
		if node.Nodes[i].isFunc() {
			decls = append(decls, node.Nodes[i].Func())
		} else if node.Nodes[i].isPackage() {
			name = ast.NewIdent(node.Nodes[i].Nodes[0].Atom)
		} else if node.Nodes[i].isImport() {
			imports.Specs = append(imports.Specs, node.Nodes[i].Import())
		}
	}

	file := &ast.File{
		Name:  name,
		Decls: decls,
		Scope: &ast.Scope{},
	}

	ast.SortImports(token.NewFileSet(), file)

	return file
}

func (node Node) Func() ast.Decl {
	var bodyList []ast.Stmt

	for _, n := range node.Nodes[3:] {
		bodyList = append(bodyList, n.Stmt())
	}

	result := &ast.FuncDecl{
		Name: ast.NewIdent(node.Nodes[0].Atom),
		Type: &ast.FuncType{
			Params:  node.Nodes[1].FieldList(),
			Results: node.Nodes[2].FieldList(),
		},
		Body: &ast.BlockStmt{
			List: bodyList,
		},
	}

	return result
}

func (node Node) FieldList() *ast.FieldList {
	var list []*ast.Field

	if node.Atom == "" {
		for i := range node.Nodes {
			list = append(list, node.Nodes[i].Field())
		}
	} else {
		list = append(list, node.Field())
	}

	return &ast.FieldList{List: list}
}

func (node Node) Field() *ast.Field {
	if len(node.Nodes) == 0 {
		return &ast.Field{Type: ast.NewIdent(node.Atom)}
	}

	return &ast.Field{Names: []*ast.Ident{ast.NewIdent(node.Atom)}, Type: ast.NewIdent(node.Nodes[0].Atom)}
}

func (node Node) Import() *ast.ImportSpec {
	return &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: node.Nodes[0].Atom,
		},
	}
}

func (node Node) Stmt() ast.Stmt {
	if node.isStmt() {
		return Stmts[node.Atom](node)
	}

	return &ast.ExprStmt{X: node.Expr()}
}

func (node Node) Expr() ast.Expr {
	if len(node.Nodes) == 0 {
		return &ast.BasicLit{Value: node.Atom}
	} else if node.isExpr() {
		return Exprs[node.Atom](node)
	} else {
		var args []ast.Expr

		for i := range node.Nodes {
			args = append(args, node.Nodes[i].Expr())
		}

		e := &ast.CallExpr{
			Fun:  ast.NewIdent(node.Atom),
			Args: args,
		}

		return e
	}
}
