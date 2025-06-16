//go:build !gui
// +build !gui

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		if arg == "--gui" || arg == "-g" {
			fmt.Println(" GUI is disabled in this build. Rebuild with: go build -tags gui")
			os.Exit(1)
		}
	}

	if len(args) == 0 {
		fmt.Println("No arguments provided. Use --help or -h for usage information.")
		os.Exit(1)
	}

	runCLI(args)
}
