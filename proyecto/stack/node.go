package stack

import (
	"edd.com/samples/structs"
)

type NodeStack struct {
	Student *structs.Student
	Next    *NodeStack
}

func (n *NodeStack) PrintNode() {
	n.Student.Print()
}
