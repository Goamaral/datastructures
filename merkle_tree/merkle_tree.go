package merkle_tree

import (
	"crypto/sha256"
	"fmt"
)

/* TREE */
type Tree struct {
	Root     *Node
	TailLeaf *Node
}

func New(data []byte) Tree {
	var tree Tree
	length := len(data)

	for offset := 0; offset < length; offset += 32 {
		upper := offset + 32
		if upper > length {
			upper = length
		}

		tree.Append(data[offset:upper])
	}

	return tree
}

func (tree *Tree) Append(data []byte) {
	tree.appendNode(NewLeafNode(data), tree.TailLeaf)
}

func (treeA Tree) Match(treeB Tree) bool {
	if treeA.Root == nil || treeB.Root == nil {
		return treeA.Root == treeB.Root
	} else {
		return treeA.Root.Checksum == treeB.Root.Checksum
	}
}

func (tree *Tree) appendNode(new *Node, tail *Node) {
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
		tree.appendNode(newParent, tail.Parent)
	}
}

/* NODE */
type Node struct {
	Checksum [sha256.Size]byte
	Parent   *Node
	Left     *Node
	Right    *Node
}

func NewParentNode(left *Node, right *Node) *Node {
	parent := &Node{Left: left, Right: right}
	left.Parent = parent

	if right == nil {
		parent.Checksum = sha256.Sum256(append(left.Checksum[:], left.Checksum[:]...))
	} else {
		right.Parent = parent
		parent.Checksum = sha256.Sum256(append(left.Checksum[:], right.Checksum[:]...))
	}

	return parent
}

func NewLeafNode(data []byte) *Node {
	return &Node{Checksum: sha256.Sum256(data)}
}

func (node *Node) UpdateChecksum() {
	node.Checksum = sha256.Sum256(append(node.Left.Checksum[:], node.Right.Checksum[:]...))
	if node.Parent != nil {
		node.Parent.UpdateChecksum()
	}
}

func (node Node) ChecksumString() string {
	return fmt.Sprintf("%x", node.Checksum)
}
