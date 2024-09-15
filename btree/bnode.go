package btree

import "fmt"


type Node struct {
	keys   []int
	t      int // branching factor
	n      int // current keys
	childs []*Node
	leaf   bool
}

// NewNode creates a new B-tree node
func NewNode(t int, leaf bool) *Node {
	return &Node{
		t:      t,
		keys:   make([]int, 0, 2*t-1),
		childs: make([]*Node, 0, 2*t),
		leaf:   leaf,
	}
}

// Traverse traverses the nodes in the B-tree
func (node *Node) traverse() {
	var i int
	for i = 0; i < node.n; i++ {
		if !node.leaf {
			node.childs[i].traverse()
		}
		fmt.Printf("%d ", node.keys[i])
	}
	if !node.leaf {
		node.childs[i].traverse()
	}
}

// InsertNonFull inserts a new key into a non-full node
func (node *Node) insertNonFull(k int) {
	i := node.n - 1

	if node.leaf {
		// If this is a leaf node, insert the key in the correct position
		node.keys = append(node.keys, 0) // Add a dummy key to extend the slice
		for i >= 0 && k < node.keys[i] {
			node.keys[i+1] = node.keys[i] // Shift keys to the right
			i--
		}
		node.keys[i+1] = k // Insert the new key
		node.n++
	} else {
		// If this is an internal node, find the child which is going to have the new key
		for i >= 0 && k < node.keys[i] {
			i--
		}
		i++
		// Check if the found child is full
		if node.childs[i].n == 2*node.t-1 {
			// If the child is full, split it
			node.splitChild(i, node.childs[i])
			// After split, the middle key of child[i] goes up and child[i] is split into two
			// See which of the two is going to have the new key
			if k > node.keys[i] {
				i++
			}
		}
		node.childs[i].insertNonFull(k) // Recur down the tree
	}
}

// SplitChild splits a full child node
func (node *Node) splitChild(i int, y *Node) {
	z := NewNode(y.t, y.leaf)
	z.n = y.t - 1

	// Copy the last (t-1) keys of y to z
	z.keys = append(z.keys, y.keys[y.t:]...)

	// Copy the last t children of y to z
	if !y.leaf {
		z.childs = append(z.childs, y.childs[y.t:]...)
	}

	// Reduce the number of keys in y
	y.n = y.t - 1

	// Insert a new child into this node
	node.childs = append(node.childs[:i+1], append([]*Node{z}, node.childs[i+1:]...)...)

	// Insert the middle key of y into this node
	node.keys = append(node.keys[:i], append([]int{y.keys[y.t-1]}, node.keys[i:]...)...)
	node.n++
}
