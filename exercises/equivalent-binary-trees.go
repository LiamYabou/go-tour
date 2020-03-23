package main

import (
	"golang.org/x/tour/tree"
	"fmt"
	"reflect"
)

func Inorder(t *tree.Tree, ch chan int) {
	if t != nil {
		Inorder(t.Left, ch)
		ch <- t.Value
		Inorder(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// We choose the algorithm of inorder traversal.
func Walk(t *tree.Tree, ch chan int) {
	Inorder(t, ch)
	close(ch)
}

// CheckSame determines whether the trees
// t1 and t2 contain the same sequence.
func CheckSame(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	t1Map, t2Map := make(map[int]int), make(map[int]int)
	for v := range ch1 {
		t1Map[v]++
	}
	for v := range ch2 {
		t2Map[v]++
	}
	return reflect.DeepEqual(t1Map, t2Map)
}

func main() {
	t1, t2 := tree.New(1), tree.New(3)
	fmt.Printf("Do t1 and t2 have the same sequence? %t", CheckSame(t1, t2))
}
