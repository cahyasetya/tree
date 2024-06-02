package binary

import (
	"fmt"
)

func InOrder(root *Node) {
	if root == nil {
		return
	}

	InOrder(root.Left)
	fmt.Println(root.Value)
	InOrder(root.Right)
}

func PreOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.Value)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func PostOrder(root *Node) {
	if root == nil {
		return
	}

	PreOrder(root.Left)
	PreOrder(root.Right)
	fmt.Println(root.Value)
}
