package loglint

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "Codestyle checks in log messages.",
	URL:  "https://github.com/accelolabs/log-lint",
	Run:  run,
	Requires: []*analysis.Analyzer{inspect.Analyzer}, 
}

func run(pass *analysis.Pass) (interface{}, error) {
	nodeFilter := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeTypes := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	nodeFilter.Preorder(nodeTypes, func(n ast.Node) {
		call := n.(*ast.CallExpr)

		if isLog(pass, call) {
			checkLogArgs(pass, call)
		}
	})

	return nil, nil
}
