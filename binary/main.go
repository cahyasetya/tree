package main

type Node struct {
	value int
	left  *Node
	right *Node
}

func main() {
	tree := &Node{
		value: 5,
		left:  &Node{value: 3, left: &Node{value: 2, left: nil, right: nil}, right: &Node{value: 4, left: nil, right: nil}},
		right: &Node{value: 7, left: &Node{value: 6, left: nil, right: nil}, right: &Node{value: 8, left: nil, right: nil}},
	}
	postOrder(tree)
}
