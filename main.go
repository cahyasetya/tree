package main

import (
	"github.com/cahyasetya/tree/binary"
	"github.com/cahyasetya/tree/nodes"
)

func main() {
	tree := &nodes.Node{
		Value: 5,
		Left:  &nodes.Node{Value: 3, Left: &nodes.Node{Value: 2, Left: nil, Right: nil}, Right: &nodes.Node{Value: 4, Left: nil, Right: nil}},
		Right: &nodes.Node{Value: 7, Left: &nodes.Node{Value: 6, Left: nil, Right: nil}, Right: &nodes.Node{Value: 8, Left: nil, Right: nil}},
	}
	binary.PostOrder(tree)
}
