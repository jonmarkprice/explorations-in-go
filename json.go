package main

import (
	"encoding/json"
	"fmt"
)

// It seems only capitalized properties are marshalled
type node struct {
	Left	*node
	Right	*node
	Leaf	int
}

func main() {
	leaf := node{nil, nil, 0}
	tree := node{
		&leaf,				// ref. external
		&node{nil, nil, 7}, // embed
		3,
	}

	enc, err := json.MarshalIndent(tree, "", "  ")

	// I think we are intended to build our objects in Go, then when finished,
	// "marshal" them to JSON.
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(enc))

	// We can also do this, which is more like JS
	anon, _ := json.Marshal(struct{Message string}{"Hello"})
	fmt.Println(string(anon))
}

