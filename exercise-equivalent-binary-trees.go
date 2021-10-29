package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var f func(t *tree.Tree, ch chan int)
	f = func(t *tree.Tree, ch chan int) {
		if t.Left != nil {
			f(t.Left, ch)
		}
		ch <- t.Value
		if t.Right != nil {
			f(t.Right, ch)
		}
	}
	f(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v := range ch1 {
		if v != <- ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	t1 := tree.New(1)
	go Walk(t1, ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
