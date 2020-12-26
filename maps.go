package main

import "fmt"

type Vertex struct {
	Lat, Lang float64
}

func main() {
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.38293829, -382.84983,
	}
	fmt.Println(m["Bell Labs"])
}
