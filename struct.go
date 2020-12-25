package main

import "fmt"

// Vertex is a point in 2D Coordinate
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
