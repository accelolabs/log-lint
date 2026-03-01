package loglint

import (
	"go/ast"
	"strings"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// inspired by https://github.com/golangci/example-plugin-module-linter

func init() {
	register.Plugin("loglint", New)
}

type AnalyzerSettings struct {
	BannedWords []string `json:"banned-words"`
}

func (s *AnalyzerSettings) Normalize() {
	for i, w := range s.BannedWords {
		s.BannedWords[i] = strings.ToLower(w)
	}
}

type LogLintPlugin struct {
	settings AnalyzerSettings
}

func New(settings any) (register.LinterPlugin, error) {
	s, err := register.DecodeSettings[AnalyzerSettings](settings)
	if err != nil {
		return nil, err
	}

	s.Normalize()

	return &LogLintPlugin{settings: s}, nil
}

func (p *LogLintPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		{
			Name:     "loglint",
			Doc:      "Codestyle checks in log messages.",
			URL:      "https://github.com/accelolabs/log-lint",
			Requires: []*analysis.Analyzer{inspect.Analyzer},
			Run:      p.run,
		},
	}, nil
}

func (p *LogLintPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

func (p *LogLintPlugin) run(pass *analysis.Pass) (any, error) {
	insp, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, nil
	}

	nodeTypes := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	insp.Preorder(nodeTypes, func(n ast.Node) {
		call := n.(*ast.CallExpr)

		if isLog(pass, call) {
			checkLogArgs(pass, call, p.settings.BannedWords)
		}
	})

	return nil, nil
}
