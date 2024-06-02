package main

import "fmt"

func inOrder(root *Node) {
	if root == nil {
		return
	}

	inOrder(root.left)
	fmt.Println(root.value)
	inOrder(root.right)
}

func preOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.value)
	preOrder(root.left)
	preOrder(root.right)
}

func postOrder(root *Node) {
	if root == nil {
		return
	}

	preOrder(root.left)
	preOrder(root.right)
	fmt.Println(root.value)
}
