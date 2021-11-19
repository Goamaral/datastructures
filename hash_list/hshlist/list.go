package hshlist

import "crypto/sha256"

/*
	Head -> Node -> Node -> Tail
*/
type hshlist struct {
	Checksum [sha256.Size]byte
	Head     *Node
	Tail     *Node
	Len      int
}

func New(data []byte) *hshlist {
	length := len(data)
	i := 0
	var list hshlist
	var hashChain []byte

	for i < length {
		upper := i + 32
		if upper > length {
			upper = length
		}

		newNode := NewNode(data[i:upper])

		if list.Len == 0 {
			list.Head = newNode
		} else {
			list.Tail.Next = newNode
		}

		list.Tail = newNode
		list.Len += 1

		hashChain = append(hashChain, newNode.Checksum[:]...)

		i += 32
	}

	list.Checksum = sha256.Sum256(hashChain)

	return &list
}

func Match(listA *hshlist, listB *hshlist) bool {
	return listA.Checksum == listB.Checksum
}
