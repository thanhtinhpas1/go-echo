package main

import "fmt"

func main() {
	pow := make([]int, 10)

	for i, _ := range pow {
		pow[i] = 1 << uint(i) // == 2 * i
	}

	for _, v := range pow {
		fmt.Printf("Value: %v\n", v)
	}
}
