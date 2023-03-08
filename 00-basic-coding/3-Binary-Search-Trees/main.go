package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// Insert a node into a binary search tree
func (node *TreeNode) Insert(val int) *TreeNode {
	if node == nil {
		return &TreeNode{val: val}
	}

	if val < node.val {
		node.left = node.left.Insert(val)
	} else {
		node.right = node.right.Insert(val)
	}

	return node
}

// Traverse a binary search tree in order
func (node *TreeNode) TraverseInOrder() {
	if node == nil {
		return
	}

	node.left.TraverseInOrder()
	fmt.Println(node.val)
	node.right.TraverseInOrder()
}

func main() {
	//@TODO: implement example of Binary Search Trees
}
