package btree


type BTree struct {
	Root *Node
	T    int // Minimum degree
}

// NewTree creates a new B-tree with a given minimum degree
func NewTree(t int) *BTree {
	return &BTree{T: t}
}

// Insert inserts a new key into the B-tree
func (t *BTree) Insert(k int) {
	if t.Root == nil {
		t.Root = NewNode(t.T, true)
		t.Root.keys = append(t.Root.keys, k)
		t.Root.n++
	} else {
		if t.Root.n == 2*t.T-1 {
			s := NewNode(t.T, false)
			s.childs = append(s.childs, t.Root)
			s.splitChild(0, t.Root)
			t.Root = s
		}
		t.Root.insertNonFull(k)
	}
}

// Traverse traverses the entire B-tree
func (t *BTree) Traverse() {
	if t.Root != nil {
		t.Root.traverse()
	}
}

// Search searches for a key in the B-tree
func (t *BTree) Search(k int) (*Node, int) {
	if t.Root == nil {
		return nil, -1
	}
	return t.Root.search(k)
}

// search searches for a key in a node
func (node *Node) search(k int) (*Node, int) {
	i := 0
	for i < node.n && k > node.keys[i] {
		i++
	}
	if i < node.n && k == node.keys[i] {
		return node, i
	}
	if node.leaf {
		return nil, -1
	}
	return node.childs[i].search(k)
}
