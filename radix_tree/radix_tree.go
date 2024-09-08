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

func (t *Tree) InsertNode(qry []byte) *Node {
	if t.Head == nil {
		t.Head = &Node{}
	}

	if qry == nil {
		t.Head.Valid = true
		return t.Head
	}

	return t.Head.InsertNode(qry)
}

func (t *Tree) SearchNode(qry []byte, exact bool) *Node {
	if t.Head == nil {
		return nil
	}

	if qry == nil {
		if t.Head.Valid {
			return t.Head
		}
		return nil
	}

	return t.Head.SearchNode(qry, exact)
}

func (t *Tree) Delete(qry []byte, exact bool) {
	n := t.SearchNode(qry, exact)
	if n != nil {
		n.Delete()
	}
}

func (n *Node) InsertNode(qry []byte) *Node {
	for _, child := range n.Children {
		prefix := extractPrefix(child.Prefix, qry)
		if len(prefix) > 0 {
			// Child prefix is a partial match -> create new node and invalidate child
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

			// Qry is an exact match -> validate child
			if len(prefix) == len(qry) {
				child.Valid = true
				return child
			}

			// Qry is a partial match -> insert in child with prefix removed
			return child.InsertNode(qry[len(prefix):])
		}
	}

	newNode := &Node{Prefix: qry, Valid: true}
	n.Children = append(n.Children, newNode)
	return newNode
}

func (n *Node) SearchNode(qry []byte, exact bool) *Node {
	for _, child := range n.Children {
		prefix := extractPrefix(child.Prefix, qry)
		if len(prefix) > 0 {
			// Exact match
			if len(child.Prefix) == len(qry) {
				return child
			}

			// Partial match -> continue search
			if len(child.Prefix) < len(qry) {
				return child.SearchNode(qry[len(child.Prefix):], exact)
			}

			if exact {
				return nil
			}
			return child
		}
	}

	return nil
}

func (n *Node) Delete() {
	n.Valid = false
	n.Children = nil
}

func extractPrefix(a, b []byte) []byte {
	var prefix []byte
	for i := 0; i < min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			break
		}
		prefix = append(prefix, a[i])
	}
	return prefix
}
