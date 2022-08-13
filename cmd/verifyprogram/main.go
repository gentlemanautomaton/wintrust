package main

import (
	"fmt"
	"os"

	"github.com/gentlemanautomaton/wintrust"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage:\n\t%s <file> [<file>...]\n", os.Args[0])
		os.Exit(1)
	}

	files := os.Args[1:]
	passed := true

	for _, file := range files {
		err := wintrust.VerifyFile(file)
		if err == nil {
			fmt.Printf("\"%s\": OK\n", file)
		} else {
			passed = false
			fmt.Printf("\"%s\": %v\n", file, err)
		}
	}

	if !passed {
		os.Exit(1)
	}
}
