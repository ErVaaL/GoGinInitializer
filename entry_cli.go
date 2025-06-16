//go:build !gui
// +build !gui

//
package main

import (
	"fmt"
	"os"
)

func runCombinedMain() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--gui" || args[0] == "-g" {
		fmt.Println("❌ GUI is disabled in this build. Rebuild with: go build -tags gui")
		os.Exit(1)
	}

	runCLI(args)
}

func runCLIMain() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--gui" || args[0] == "-g" {
		fmt.Println("❌ GUI is disabled in this build. Rebuild with: go build -tags gui")
		os.Exit(1)
	}

	runCLI(args)
}
