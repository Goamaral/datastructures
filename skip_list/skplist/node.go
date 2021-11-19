package skplist

// import "strconv"
import "fmt"

type Node struct {
	Value int
	Next  *Node
	Child *Node
}

func NewNode(value int) *Node {
	newNode := &Node{Value: value}
	return newNode
}

func (node *Node) Tail() *Node {
	current := node
	for ; current.Next != nil; current = current.Next {
	}
	return current
}

func (node *Node) ToString() string {
	return fmt.Sprintf("%d", node.Value)
}
