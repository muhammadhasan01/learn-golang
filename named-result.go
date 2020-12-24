package main

import "fmt"

func split(num int) (x, y int) {
	x = num / 10
	y = num % 10
	return
}

func main() {
	fmt.Println(split(15))
}
