package main

import "fmt"

func main() {
	// Slice literal is like an array literal without the lengh
	q := []int{2, 3, 1, 5, 2, 5}
	fmt.Print(q)

	r := []bool{true, false, true, false, false, false}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}

	fmt.Println(s)
}
