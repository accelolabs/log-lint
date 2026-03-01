package main

import (
	loglint "github.com/accelolabs/log-lint"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(loglint.Analyzer)
}
