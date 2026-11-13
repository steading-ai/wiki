// Command wikilint checks the wiki against the rules in AGENTS.md and prints
// each problem as path:line: message. It exits 1 when it finds any, and 2 when
// it cannot run.
//
//	wikilint [root]
//
// root is the top of the wiki, the current directory by default.
package main

import (
	"fmt"
	"os"

	"github.com/steading-ai/wiki/internal/lint"
)

func main() {
	root := "."
	switch len(os.Args) {
	case 1:
	case 2:
		root = os.Args[1]
	default:
		fmt.Fprintln(os.Stderr, "usage: wikilint [root]")
		os.Exit(2)
	}
	problems, err := lint.Check(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, "wikilint:", err)
		os.Exit(2)
	}
	for _, p := range problems {
		fmt.Println(p)
	}
	if len(problems) > 0 {
		os.Exit(1)
	}
}
