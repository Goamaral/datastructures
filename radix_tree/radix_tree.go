package radix_tree

type Tree struct {
	Head *Node
}

type Node struct {
	Prefix   []byte
	Valid    bool
	Children []*Node
}

func New() Tree {
	return Tree{}
}

func (t *Tree) Insert(qry []byte) *Node {
	if t.Head == nil {
		t.Head = &Node{}
	}

	if qry == nil {
		t.Head.Valid = true
		return t.Head
	}

	return t.Head.Insert(qry)
}

func (t *Tree) ExactSearch(qry []byte) *Node {
	if t.Head == nil {
		return nil
	}

	if qry == nil {
		if t.Head.Valid {
			return t.Head
		}
		return nil
	}

	return t.Head.ExactSearch(qry)
}

// func (t *Tree) PrefixSearch(qry []byte) []*Node {
// }

// func (t *Tree) Delete(qry string) {
// }

func (n *Node) Insert(qry []byte) *Node {
	for _, child := range n.Children {
		prefix := child.extractPrefix(qry)
		if len(prefix) > 0 {
			if len(prefix) == len(qry) {
				child.Valid = true
				return child

			} else {
				// Child prefix is bigger than prefix -> create new node and invalidate child
				if len(child.Prefix) > len(prefix) {
					left := &Node{
						Prefix:   make([]byte, len(child.Prefix)-len(prefix)),
						Valid:    child.Valid,
						Children: child.Children,
					}
					copy(left.Prefix, child.Prefix[len(prefix):])
					child.Valid = false
					child.Children = []*Node{left}
					child.Prefix = prefix
				}

				// Qry matches prefix -> qry is the child prefix
				if len(prefix) == len(qry) {
					child.Valid = true
					return child

					// Qry is bigger than prefix -> insert in child with prefix removed
				} else {
					return child.Insert(qry[len(prefix):])
				}
			}
		}
	}

	newNode := &Node{Prefix: qry, Valid: true}
	n.Children = append(n.Children, newNode)
	return newNode
}

func (n *Node) ExactSearch(qry []byte) *Node {
	for _, child := range n.Children {
		prefix := child.extractPrefix(qry)
		if len(prefix) > 0 {
			if len(prefix) == len(qry) {
				return child

			} else {
				if len(child.Prefix) == len(qry) {
					return child
				} else {
					return child.ExactSearch(qry[len(prefix):])
				}
			}
		}
	}

	return nil
}

func (n Node) extractPrefix(qry []byte) []byte {
	var prefix []byte
	for i := 0; i < min(len(n.Prefix), len(qry)); i++ {
		if n.Prefix[i] != qry[i] {
			break
		}
		prefix = append(prefix, n.Prefix[i])
	}
	return prefix
}
