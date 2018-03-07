package main

import "fmt"

func f(xs []int) {
	fmt.Printf("[ ")
	for _, x := range xs {
		fmt.Printf("%d, ", x)
		x++
	}
	fmt.Println(" ]")
}

func main() {
	f([]int{1, 2, 3}) // slice

	// f([...]int{1, 2, 3}) // No

	a := [...]int{1, 2, 3}
	f(a[:]) // Take a slice of an array (by copy)

	f(a[:0])
	f(a[:2])
}
