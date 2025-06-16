//go:build gui
// +build gui

package main

import (
	"os"
)

func runCombinedMain() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--gui" || args[0] == "-g" {
		launchGui()
		return
	}

	runCLI(args)
}
