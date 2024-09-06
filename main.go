package main

import (
	binarytree "github.com/cahyasetya/tree/binarytree"
	"github.com/cahyasetya/tree/iterator"
	"github.com/cahyasetya/tree/model"
)

func main() {
	root := &model.Node{
		Key: 2,
	}
	binarytree.Insert(root, 1)
	binarytree.Insert(root, 3)
	iterator.LevelOrder(root)
}
