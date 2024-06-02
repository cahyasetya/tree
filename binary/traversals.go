package binary

import (
	"fmt"

	"github.com/cahyasetya/tree/nodes"
)

func InOrder(root *nodes.Node) {
	if root == nil {
		return
	}

	InOrder(root.Left)
	fmt.Println(root.Value)
	InOrder(root.Right)
}

func PreOrder(root *nodes.Node) {
	if root == nil {
		return
	}

	fmt.Println(root.Value)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func PostOrder(root *nodes.Node) {
	if root == nil {
		return
	}

	PreOrder(root.Left)
	PreOrder(root.Right)
	fmt.Println(root.Value)
}
