package main

import "fmt"

// Go has only one looping construct, the for loop
func main() {
	// 3 parts
	// init statement
	// condition
	// post execute
	// Note: The init and post statements are optional
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	// While looping
	for 1 == 1 {
		fmt.Println("End while")
		break
	}

	mapStr := make([]int, 3, 3)
	for i1, i2 := range mapStr {
		fmt.Println(i1)
		fmt.Println(i2)
	}

	fmt.Println(sum)
}
