package main

import "fmt"

func main() {
	// An array has a fixed size
	// Slice is dynamically-sized
	// A slice is formed by specifying two indices, a low high bound, seperated by a colon

	primes := [6]int{2, 3, 1, 6, 2, 6}
	fmt.Println(primes)
}
