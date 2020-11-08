package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func Walker(t *tree.Tree, ch chan int){
	Walk(t, ch)
	close(ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if t == nil {
		return
	}else if t.Left == nil {
		ch <- t.Value
		if t.Right != nil {
			Walk(t.Right, ch)
		}
		return
	}else{
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
			Walk(t.Right, ch)
		}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walker(t1, ch1)
	go Walker(t2, ch2)
	
	flag := true
	
	for i := range ch1{
		if i == <- ch2 {
			fmt.Println(i)
			flag = true
		} else {
			flag = false
		}
	}
	return flag
}

func main() {
	ch := make(chan int)
	
	t1 := tree.New(1)
	fmt.Println(t1)
	go Walker(t1, ch)
	for i := range ch {
		fmt.Println(i)
	}
	t2 := tree.New(1)
	//fmt.Println(t2)
	fmt.Println(Same(t1,t2))
	
}
