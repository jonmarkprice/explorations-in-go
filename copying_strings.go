package main

import "fmt"

func main() {
	// Strings in Go are not pointers, like in C.
	// Assignment copies, rather than referencing.

	s1 := "Hello"	// create string
	s2 := s1		// implicit copy
	s3 := &s1		// explict reference

	s1 += " world"	// mutate original

	fmt.Println("1. ", s1)
	fmt.Println("2. ", s2)
	fmt.Println("3. ", *s3)	// must deref.
}
