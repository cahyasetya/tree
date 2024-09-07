package search

import "github.com/cahyasetya/tree/model"

func Search(node *model.Node, key int) *model.Node {
	if node == nil {
		return nil
	}

	if key == node.Key {
		return node
	}
	if key < node.Key {
		return Search(node.Left, key)
	}
	if key > node.Key {
		return Search(node.Right, key)
	}
	return nil
}
