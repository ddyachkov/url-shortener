package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestMainAnazyler(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), MainAnalyzer, "./...")
}
