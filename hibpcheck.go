package main

import (
	"fmt"
	"os"

	"github.com/iand/hibp"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: hibpcheck <password>\n")
		os.Exit(1)
	}

	count, err := hibp.Lookup(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed lookup: %v\n", err)
		os.Exit(1)
	}
	println(count)
}
