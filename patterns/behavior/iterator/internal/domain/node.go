package domain

type Node struct {
	Value               int
	Left, Right, Parent *Node
}

func NewNode(value int, left *Node, right *Node) *Node {
	n := &Node{Value: value, Left: left, Right: right}
	left.Parent = n
	right.Parent = n
	return n
}

func NewTerminalNode(value int) *Node {
	return &Node{Value: value}
}
