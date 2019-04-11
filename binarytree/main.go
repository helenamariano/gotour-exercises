package main

import (
	"fmt"
	"math/rand"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan<- int, stop <-chan bool) {
	doWalk(t, ch, stop)
	close(ch)
}

func doWalk(t *tree.Tree, ch chan<- int, stop <-chan bool) {
	if t.Left != nil {
		doWalk(t.Left, ch, stop)
	}
	select {
	case ch <- t.Value:
	case <-stop:
		return
	}
	if t.Right != nil {
		doWalk(t.Right, ch, stop)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	stop := make(chan bool)
	defer close(stop)

	go Walk(t1, ch1, stop)
	go Walk(t2, ch2, stop)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		// one tree has more nodes than the other
		if ok1 != ok2 {
			return false
		}

		if !ok1 {
			return true
		}

		if v1 != v2 {
			return false
		}
	}
}

func main() {
	testWalk()
	testSameEquivalentTrees()
	testSameDifferentTrees()
	testSameDifferentTrees2()
}

func testWalk() {
	fmt.Println("---------")
	fmt.Println("Test walk")
	fmt.Println("---------")

	ch := make(chan int)
	t := tree.New(1)
	fmt.Println(t)
	go Walk(t, ch, nil)
	for i := range ch {
		fmt.Println(i)
	}
}

func testSameEquivalentTrees() {
	fmt.Println("----------------------------")
	fmt.Println("Test same (equivalent trees)")
	fmt.Println("----------------------------")

	t1, t2 := tree.New(1), tree.New(1)
	fmt.Printf("Tree 1: %s\n", t1.String())
	fmt.Printf("Tree 2: %s\n", t2.String())
	same := Same(t1, t2)
	fmt.Printf("Same: %t\n", same)
}

func testSameDifferentTrees() {
	fmt.Println("--------------------------------")
	fmt.Println("Test same (different node values)")
	fmt.Println("--------------------------------")

	t1, t2 := tree.New(1), tree.New(20)
	fmt.Printf("Tree 1: %s\n", t1.String())
	fmt.Printf("Tree 2: %s\n", t2.String())
	same := Same(t1, t2)
	fmt.Printf("Same: %t\n", same)
}

// the tree constructor from the golang.org/x/tour/tree always creates
// the 10 nodes for each tree but I want to test that Same() works
// even if the number of nodes between the tress are different
func testSameDifferentTrees2() {
	fmt.Println("--------------------------------------")
	fmt.Println("Test same (different number of nodes)")
	fmt.Println("--------------------------------------")

	t1, t2 := NewTreeVarNodes(1, 10), NewTreeVarNodes(1, 100)
	fmt.Printf("Tree 1: %s\n", t1.String())
	fmt.Printf("Tree 2: %s\n", t2.String())
	same := Same(t1, t2)
	fmt.Printf("Same: %t\n", same)
}

// NewTreeVarNodes returns a new, random binary tree with n nodes holding the values k, 2k, ..., nk.
func NewTreeVarNodes(k, n int) *tree.Tree {
	var t *tree.Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *tree.Tree, v int) *tree.Tree {
	if t == nil {
		return &tree.Tree{Left: nil, Value: v, Right: nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}
