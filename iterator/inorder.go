package iterator

import (
	"fmt"
	"github.com/cahyasetya/tree/model"
)

func Inorder(root *model.Node) {
	if root == nil {
		return
	}

	Inorder(root.Left)
	fmt.Println(root.Key)
	Inorder(root.Right)
}

func Preorder(root *model.Node) {
	if root == nil {
		return
	}

	fmt.Println(root.Key)
	Preorder(root.Left)
	Preorder(root.Right)
}

func Postorder(root *model.Node) {
	if root == nil {
		return
	}

	Postorder(root.Left)
	Postorder(root.Right)
	fmt.Println(root.Key)
}

func LevelOrder(root *model.Node) {
	if root == nil {
		return
	}

	queue := []*model.Node{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println(node.Key)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
