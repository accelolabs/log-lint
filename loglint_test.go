package loglint

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLogLint(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "testdata/loglint")
}
