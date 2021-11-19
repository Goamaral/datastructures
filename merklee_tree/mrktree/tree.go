package mrktree

type mrktree struct {
	Root     *Node
	TailLeaf *Node
}

func New(data []byte) *mrktree {
	length := len(data)
	i := 0
	var tree mrktree

	for i < length {
		upper := i + 32
		if upper > length {
			upper = length
		}

		newLeaf := NewLeafNode(data[i:upper])
		tree.AppendNode(newLeaf, tree.TailLeaf)

		i += 32
	}

	return &tree
}

func (tree *mrktree) AppendNode(new *Node, tail *Node) {
	if tail == nil {
		tree.Root = new
		tree.TailLeaf = new
	} else if tail.Parent == nil {
		tree.Root = NewParentNode(tail, new)
	} else if tail.Parent.Right == nil {
		tail.Parent.Right = new
		tail.Parent.UpdateChecksum()
	} else {
		newParent := NewParentNode(new, nil)
		tree.AppendNode(newParent, tail.Parent)
	}
}

func Match(treeA *mrktree, treeB *mrktree) bool {
	if treeA.Root == nil || treeB.Root == nil {
		return treeA.Root == treeB.Root
	} else {
		return treeA.Root.Checksum == treeB.Root.Checksum
	}
}
