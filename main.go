package main

import (
	"github.com/sergei-dyshel/fzf-abbrev/src"
	"github.com/sergei-dyshel/fzf-abbrev/src/protector"
)

var revision string

func main() {
	protector.Protect()
	fzf.Run(fzf.ParseOptions(), revision)
}
