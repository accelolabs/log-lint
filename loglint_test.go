package loglint

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLogLint(t *testing.T) {
	rawSettings := map[string]any{
		"banned-words": []string{"password", "apikey", "token"},
	}

	plugin, err := New(rawSettings)
	if err != nil {
		t.Fatalf("failed to create plugin: %v", err)
	}

	analyzers, err := plugin.BuildAnalyzers()
	if err != nil || len(analyzers) == 0 {
		t.Fatalf("failed to build analyzers: %v", err)
	}

	analysistest.Run(t, analysistest.TestData(), analyzers[0], "testdata/loglint")
}
