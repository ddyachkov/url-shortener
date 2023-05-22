package main

import (
	"strings"

	"github.com/ddyachkov/url-shortener/internal/analyzer"
	"github.com/kisielk/errcheck/errcheck"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
)

func main() {
	myChecks := []*analysis.Analyzer{
		analyzer.MainAnalyzer,
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
		errcheck.Analyzer,
	}

	stChecks := map[string]bool{
		"S1000":  true,
		"ST1005": true,
		"QF1002": true,
	}
	for _, v := range staticcheck.Analyzers {
		if strings.HasPrefix(v.Analyzer.Name, "SA") || stChecks[v.Analyzer.Name] {
			myChecks = append(myChecks, v.Analyzer)
		}
	}
	multichecker.Main(
		myChecks...,
	)

}
