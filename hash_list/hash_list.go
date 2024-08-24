package hash_list

import "crypto/sha256"

/* LIST */
// Head -> Node -> Node
type List struct {
	Checksum [sha256.Size]byte
	Head     *Node
	Len      int
}

func New(data []byte) List {
	var list List
	var hashChain []byte
	length := len(data)

	for offset := 0; offset < length; offset += 32 {
		upper := offset + 32
		if upper > length {
			upper = length
		}

		newNode := NewNode(data[offset:upper])

		if list.Len == 0 {
			list.Head = newNode
		}
		list.Len++

		hashChain = append(hashChain, newNode.Checksum[:]...)
	}

	list.Checksum = sha256.Sum256(hashChain)

	return list
}

func (listA List) Match(listB List) bool {
	return listA.Checksum == listB.Checksum
}

/* NODE */
type Node struct {
	Checksum [sha256.Size]byte
	Next     *Node
}

func NewNode(data []byte) *Node {
	return &Node{Checksum: sha256.Sum256(data)}
}
