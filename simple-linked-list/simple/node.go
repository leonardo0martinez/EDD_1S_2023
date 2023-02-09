package simple

//El nodo solito xd
type Node struct {
	value int
	next  *Node
}

func (n *Node) SetValue(value int) {
	n.value = value
}

func (n *Node) GetValue() int {
	return n.value
}
