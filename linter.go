package loglint

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
	"unicode"

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
	return path == "log/slog" || path == "go.uber.org/zap"
}

func checkLogArgs(pass *analysis.Pass, call *ast.CallExpr, bannedWords []string) {
	for _, arg := range call.Args {
		ast.Inspect(arg, func(n ast.Node) bool {
			switch argType := n.(type) {
			case *ast.Ident:
				name := strings.ToLower(argType.Name)

				for _, bannedWord := range bannedWords {
					if strings.Contains(strings.ToLower(name), bannedWord) {
						pass.Reportf(argType.Pos(), "log check failed: message should not contain potential secrets")
						break
					}
				}

			case *ast.BasicLit:
				if argType.Kind != token.STRING {
					return true
				}

				val, err := strconv.Unquote(argType.Value)
				if err != nil || val == "" {
					return true
				}

				if errMsg := checkLiteralRules(val); errMsg != "" {
					pass.Reportf(argType.Pos(), "log check failed: %s", errMsg)
				}
			}
			return true
		})
	}
}

func checkLiteralRules(s string) string {
	for _, r := range s {
		if unicode.IsSpace(r) {
			continue
		}
		if unicode.IsUpper(r) {
			return "message should start with a lowercase letter"
		}
		break
	}

	for _, r := range s {
		if !(unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r)) {
			return "message should not contain special symbols"
		}

		if r > unicode.MaxASCII {
			return "message should be in english"
		}
	}

	return ""
}
