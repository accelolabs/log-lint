package loglint

import (
	"go/ast"
	"go/token"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
)

func isLog(pass *analysis.Pass, call *ast.CallExpr) bool {
	selector, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	pkgObject := pass.TypesInfo.Uses[selector.Sel]
	if pkgObject == nil {
		return false
	}

	pkg := pkgObject.Pkg()
	if pkg == nil {
		return false
	}

	name := selector.Sel.Name
	switch name {
	case "Debug", "Info", "Warn", "Error", "Fatal", "Panic",
		"Debugf", "Infof", "Warnf", "Errorf", "Fatalf", "Panicf":
	default:
		return false
	}

	path := pkg.Path()
	return path == "log/slog" || strings.Contains(path, "go.uber.org/zap")
}

func checkLogArgs(pass *analysis.Pass, call *ast.CallExpr, bannedWords []string) {
	for _, arg := range call.Args {
		ast.Inspect(arg, func(n ast.Node) bool {
			ident, ok := n.(*ast.Ident)
			if !ok {
				return true
			}

			for _, bannedWord := range bannedWords {
				if strings.Contains(strings.ToLower(ident.Name), bannedWord) {
					pass.Reportf(ident.Pos(), "log check failed: message should not contain potential secrets")
					break
				}
			}
			return true
		})

		ast.Inspect(arg, func(n ast.Node) bool {
			lit, ok := n.(*ast.BasicLit)
			if !ok || lit.Kind != token.STRING {
				return true
			}

			val := strings.Trim(lit.Value, `"`+"`")
			if val == "" {
				return true
			}

			if errMsg := checkLiteralRules(val); errMsg != "" {
				pass.Reportf(lit.Pos(), "log check failed: %s", errMsg)
			}
			return true
		})
	}
}

func checkLiteralRules(s string) string {
	firstRune, _ := utf8.DecodeRuneInString(s)
	if unicode.IsUpper(firstRune) {
		return "message should start with a lowercase letter"
	}

	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && !unicode.IsSpace(r) {
			return "message should not contain special symbols"
		}

		if r > unicode.MaxASCII {
			return "message should be in english"
		}
	}

	return ""
}
