package main

import (
	"fmt"
	loglint "github.com/accelolabs/log-lint"

	"golang.org/x/tools/go/analysis"
)

// https://github.com/golangci/example-plugin-linter
func New(conf any) ([]*analysis.Analyzer, error) {
	fmt.Printf("Configuration (%[1]T): %#[1]v\n", conf)
	return []*analysis.Analyzer{loglint.Analyzer}, nil
}
