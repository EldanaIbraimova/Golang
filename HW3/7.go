package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func Walk(t *tree.Tree, ch chan int){
	if t == nil { return }
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch<-t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1,ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i:= 0; i<10; i++ {
		if <-ch1 != <-ch2 {
			return false;
		}
	}
	return true
}

func main() {
	size:=10
	ch := make(chan int, size)
	go Walk(tree.New(1),ch)
	for i:= 0; i < size; i++ {
		fmt.Print(<-ch)
		fmt.Print(" ")
	}
	fmt.Println()
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
