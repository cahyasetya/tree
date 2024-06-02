package binary

import (
	"container/list"
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

func LevelOrder(root *Node) {
	lst := list.New()
	lst.PushBack(root)

	for lst.Len() != 0 {
		front := lst.Front().Value.(*Node)
		if front != nil {
			fmt.Println(front.Value)
			if front.Left != nil {
				lst.PushBack(front.Left)
			}
			if front.Right != nil {
				lst.PushBack(front.Right)
			}
			lst.Remove(lst.Front())
		}
	}
}
