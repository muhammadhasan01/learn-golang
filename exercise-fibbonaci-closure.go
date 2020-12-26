package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib := make([]int, 10)
	fib[1] = 1
	for i := 2; i < 10; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	it := -1
	return func() int {
		it++
		return fib[it]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
