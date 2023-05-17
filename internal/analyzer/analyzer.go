package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// MainAnalyzer analyzes main function in main packages
var MainAnalyzer = &analysis.Analyzer{
	Name: "maincheck",
	Doc:  "checks for main function in main package",
	Run:  mainAnalyzerRun,
}

func mainAnalyzerRun(pass *analysis.Pass) (interface{}, error) {
	osExitCheck := func(se *ast.SelectorExpr) {
		ident, ok := se.X.(*ast.Ident)
		if !ok {
			return
		}
		if ident.Name == "os" && se.Sel.Name == "Exit" {
			pass.Reportf(ident.Pos(), "cannot call os.Exit in function main in package main")
		}
	}

	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.File:
				if x.Name.Name != "main" {
					return false
				}
			case *ast.FuncDecl:
				if x.Name.Name != "main" {
					return false
				}
			case *ast.SelectorExpr:
				osExitCheck(x)
			}
			return true
		})
	}
	return nil, nil
}
