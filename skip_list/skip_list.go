package skip_list

import "fmt"

/*
Example height 4, skips 2
1                 -              15
1     -     -     10       -     15
1 - - 4 - - 7 - - 10 -  -  13 -  15
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
*/
type List struct {
	skips      int
	height     int
	levelRoots []*Node
}

// height >= 1
// skips >= 0
func New(height int, skips int) *List {
	list := List{skips: skips, height: height}
	list.levelRoots = make([]*Node, height)
	return &list
}

// Insert new value, in order, at floor level.
// Then update all upper levels.
func (list *List) Insert(value int) {
	newNode := NewNode(value)
	levelRoot := list.levelRoots[0]

	if levelRoot == nil || levelRoot.Value > value {
		list.levelRoots[0] = newNode
		newNode.Next = levelRoot
	} else {
		// TODO: Use Search method
		current := levelRoot
		for ; current.Next != nil && current.Next.Value <= value; current = current.Next {
		}

		oldNext := current.Next
		current.Next = newNode
		newNode.Next = oldNext
	}

	for i := 1; i < list.height; i++ {
		list.updateLevel(i)
	}
}

// Updates level nodes based on the previous level
// Nodes are added/updated after n skiped or if the lowerLevelNode is the last
func (list *List) updateLevel(level int) {
	lowerLevelNode := list.levelRoots[level-1]
	levelNode := list.levelRoots[level]
	prevLevelNode := list.levelRoots[level]
	i := 0

	for ; lowerLevelNode != nil; lowerLevelNode = lowerLevelNode.Next {
		if i%(list.skips+1) == 0 || lowerLevelNode.Next == nil {

			// If current level empty, initialize level
			if list.levelRoots[level] == nil {
				prevLevelNode = NewNode(lowerLevelNode.Value)
				prevLevelNode.Child = lowerLevelNode
				list.levelRoots[level] = prevLevelNode
			} else {
				// If the levelNode is empty, initialize and append it to the current level
				if levelNode == nil {
					levelNode = NewNode(lowerLevelNode.Value)
					levelNode.Child = lowerLevelNode
					prevLevelNode.Next = levelNode
					prevLevelNode = levelNode
					// If the levelNode exists, update its data (value, child)
				} else {
					levelNode.Value = lowerLevelNode.Value
					levelNode.Child = lowerLevelNode
					prevLevelNode = levelNode
					levelNode = levelNode.Next
				}
			}
		}

		i++
	}
}

// Search target value
func (list *List) Search(value int) bool {
	return list.search(value, list.levelRoots[list.height-1])
}

// Print nodes organized by level
func (list *List) PrintLevels() {
	for index := range list.levelRoots {
		level := list.height - index - 1
		list.PrintLevel(level)
	}
}

// Print level nodes
func (list *List) PrintLevel(level int) {
	levelRoot := list.levelRoots[level]

	fmt.Print("Level[", level, "] ")
	for current := levelRoot; current != nil; current = current.Next {
		fmt.Print(current.ToString(), " ")
	}
	fmt.Print("\n")
}

// Print level nodes from level root
func (list *List) PrintLevelFromRoot(levelRoot *Node) {
	fmt.Print("Level ")
	for current := levelRoot; current != nil; current = current.Next {
		fmt.Print(current.ToString(), " ")
	}
	fmt.Print("\n")
}

// Search target value from levelNode onwards
// Recur to lower levels until floor level is reached
func (list *List) search(value int, levelNode *Node) bool {
	// If the node does not exist or if the value is less than the first value
	if levelNode == nil || value < levelNode.Value {
		return false
	} else {
		prevLevelNode := levelNode

		for levelNode = prevLevelNode.Next; !(levelNode == nil || value < levelNode.Value); levelNode = levelNode.Next {
			prevLevelNode = levelNode
		}

		// If we are at floor level
		if prevLevelNode.Child == nil {
			return prevLevelNode.Value == value
		} else {
			return list.search(value, prevLevelNode.Child)
		}
	}
}

/* NODE */
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
