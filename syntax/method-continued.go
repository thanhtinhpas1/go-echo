package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// You can declare a method on non-struct types, too
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main()  {
	f := MyFloat(-math.Sqrt(12))
	fmt.Println(f.Abs())
}