package main

import (
	"fmt"
	"flag"
)

func main() {
	// Make an empty slice, only initialize it based on a condition.
	// This can be clearer than repeatedly calling append.
	s := make([]uint, 0)

	init := flag.Bool("init", false, "intialize the list")
	flag.Parse()

	if *init {
		s = []uint{1, 1, 2, 5, 14, 42}
	}

	fmt.Printf("The first %d Catalan numbers are %v: ", len(s), s)
}
