package main

import (
	"fmt"
	"container/list"
)

type MyStack struct {
	List *list.List
}

type BinaryTree struct {
	Value interface{}
	Left *BinaryTree
	Right *BinaryTree
}

type Tree struct {
	Value interface{}
	Children []*Tree
}

func (stack *MyStack) pop()interface{} {
	if elem := stack.List.Back(); elem != nil {
		stack.List.Remove(elem)
		return elem.Value
	}
	return nil
}

func (stack *MyStack) push(elem interface{}) {
	stack.List.PushBack(elem)
}

func preOrderRecur(node *BinaryTree) {
	if node == nil {
		return
	}

	fmt.Println(node.Value)
	preOrderRecur(node.Left)
	preOrderRecur(node.Right)
}

func main() {

	node7 := &BinaryTree{Value: 7}
	node6 := &BinaryTree{Value: 6}
	node5 := &BinaryTree{Value: 5}
	node4 := &BinaryTree{Value: 4}
	node3 := &BinaryTree{Value: 3, Left: node6, Right: node7}
	node2 := &BinaryTree{Value: 2, Left: node4, Right: node5}
	root := &BinaryTree{Value: 1, Left: node2, Right: node3}
	preOrderRecur(root)
	fmt.Println()
	//preOrder(root)

}
