package main

import "fmt"

var pow = []int{1, 2, 4, 6, 8, 10, 12, 14}

func main() {
	for index, value := range pow {
		fmt.Printf("2**%d = %d\n", index, value)
	}
}
