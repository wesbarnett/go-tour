package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	for _, v := range strings.Fields(s) {
		result[v] += 1
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
