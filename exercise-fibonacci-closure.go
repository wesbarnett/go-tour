package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f1 := 0
	f2 := 1
	f := 0
	i := -1
	return func() int {
		f = f1 + f2
		f2 = f1
		f1 = f
		if i <= 0 {
			i += 1
			return i
		} else {
			return f
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
