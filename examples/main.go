package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		basic()
		return
	}
	switch os.Args[1] {
	case "basic":
		basic()
	case "namespace":
		namespace()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\ncommand must be basic or namespace", os.Args[1])
		os.Exit(1)
	}
}
