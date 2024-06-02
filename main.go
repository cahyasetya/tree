package main

import (
	"github.com/cahyasetya/tree/binary"
)

func main() {
	tree := &binary.Node{
		Value: 5,
		Left:  &binary.Node{Value: 3, Left: &binary.Node{Value: 2, Left: nil, Right: nil}, Right: &binary.Node{Value: 4, Left: nil, Right: nil}},
		Right: &binary.Node{Value: 7, Left: &binary.Node{Value: 6, Left: nil, Right: nil}, Right: &binary.Node{Value: 8, Left: nil, Right: nil}},
	}
	binary.LevelOrder(tree)
}
