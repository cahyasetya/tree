package avl

import "github.com/cahyasetya/tree/model"

func getHeight(node *model.Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func rightRotate(node *model.Node) *model.Node {
	if node == nil {
		return nil
	}

	y := node.Left
	T2 := y.Right

	y.Right = node
	node.Left = T2

	node.Height = 1 + max(getHeight(node.Left), getHeight(node.Right))
	y.Height = 1 + max(getHeight(y.Left), getHeight(y.Right))
	return y
}

func leftRotate(node *model.Node) *model.Node {
	if node == nil {
		return nil
	}

	y := node.Right
	T2 := y.Left

	y.Left = node
	node.Right = T2

	node.Height = 1 + max(getHeight(node.Left), getHeight(node.Right))
	y.Height = 1 + max(getHeight(y.Left), getHeight(y.Right))

	return y
}

func getBalance(node *model.Node) int {
	if node == nil {
		return 0
	}

	leftHeight := getHeight(node.Left)
	rightHeight := getHeight(node.Right)

	return leftHeight - rightHeight
}

func Insert(node *model.Node, key int) *model.Node {
	if node == nil {
		return newNode(key)
	}

	if key < node.Key {
		node.Left = Insert(node.Left, key)
	} else if key > node.Key {
		node.Right = Insert(node.Right, key)
	} else {
		return node
	}

	node.Height = 1 + max(getHeight(node.Left), getHeight(node.Right))

	balance := getBalance(node)

	if balance > 1 && key < node.Left.Key {
		return rightRotate(node)
	}

	if balance < -1 && key > node.Right.Key {
		return leftRotate(node)
	}

	if balance > 1 && key > node.Left.Key {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	if balance < -1 && key < node.Right.Key {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

func newNode(key int) *model.Node {
	return &model.Node{
		Key:    key,
		Height: 1,
	}
}
