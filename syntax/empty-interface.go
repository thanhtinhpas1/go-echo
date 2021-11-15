package main

import "fmt"

func main() {
	// The interface type that specifies with zero methods is known as the `empty interface`
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}