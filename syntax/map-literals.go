package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.59888, -34.2993},
	"Google": {
		38.59888, -29.29832,
	},
}

func main() {
	fmt.Println(m)
}
