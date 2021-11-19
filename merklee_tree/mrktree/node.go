package mrktree

import (
	"crypto/sha256"
	"fmt"
)

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
	return &Node{Checksum: sha256.Sum256(data[:32])}
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
