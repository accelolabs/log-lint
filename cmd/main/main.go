package main

import (
	loglint "github.com/accelolabs/log-lint"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	rawSettings := map[string]any{
		"banned-words": []string{"password", "apikey", "token"},
	}

	plugin, err := loglint.New(rawSettings)
	if err != nil {
		panic(err)
	}

	analyzers, err := plugin.BuildAnalyzers()
	if err != nil || len(analyzers) == 0 {
		panic("no analyzers found")
	}

	singlechecker.Main(analyzers[0])
}
