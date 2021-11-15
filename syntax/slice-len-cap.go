package main

import "fmt"

func main() {
	// A slice has both length can capacity
	// The zero value of slice is `Nil`
	s := []int{2, 3, 2, 1, 6}
	printSlice(s)
}

func printSlice(s []int)  {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}