package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	words := strings.Split(s, " ")
	for _, v := range words {
		ret[v] = ret[v] + 1
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
