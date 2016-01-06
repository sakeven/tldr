package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage: tldr <cmd>")
}

func main() {
	if len(os.Args) != 2 {
		usage()
		return
	}

	cmd := os.Args[1]
	rd, err := getTldr("common", cmd)
	if err != nil {
		os.Exit(1)
	}

	Render(rd)
}
