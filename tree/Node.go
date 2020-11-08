package main

import "fmt"

type TreeNode struct {
	value       int
	left, right *TreeNode
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

// receiver declared. Create a function of struct TreeNode
func (node TreeNode) Print() {
	fmt.Println(node.value)

}

//traverse a tree - Middle
func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.left.Traverse()
	node.Print()
	node.right.Traverse()

}

func (node *TreeNode) SetValue(v int) {
	node.value = v
}
func main() {

	// create a binary tree
	root := TreeNode{value: 3}

	root.left = &TreeNode{value: 0}
	root.left.right = &TreeNode{value: 2}
	root.right = &TreeNode{value: 5}
	root.right.left = &TreeNode{value: 0}
	fmt.Println(root)
	fmt.Println(root.right.value)
	root.Print()
	root.right.left.SetValue(5555)

	// traversing the tree
	fmt.Println("start traversing")
	root.Traverse()
}
