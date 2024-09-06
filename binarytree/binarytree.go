package binarytree

import (
	"github.com/cahyasetya/tree/model"
)

func Insert(node *model.Node, key int) {
	if node == nil {
		return
	}

	if key <= node.Key {
		if node.Left == nil {
			node.Left = newNode(key)
		} else {
			Insert(node.Left, key)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode(key)
		} else {
			Insert(node.Right, key)
		}
	}
}

func newNode(key int) *model.Node {
	return &model.Node{
		Key: key,
	}
}
