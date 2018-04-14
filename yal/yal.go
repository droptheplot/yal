package yal

import (
	goast "go/ast"

	"github.com/droptheplot/yal/yal/ast"
	"github.com/droptheplot/yal/yal/core"
)

func New(src []byte) *goast.File {
	node := ast.New(src)

	return core.File(node)
}
