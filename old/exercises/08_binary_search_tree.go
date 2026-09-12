package main

import "fmt"

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}

	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}

	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

type BSTAdder struct {
	start int
}

func (a BSTAdder) AddTo(val int) int {
	return a.start + val
}

func Run08BinarySearchTree() {
	var it *IntTree

	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(9)
	it = it.Insert(13)

	fmt.Println(it.Contains(7))

	a := BSTAdder{start: 10}
	fmt.Println(a.AddTo(20))
}
