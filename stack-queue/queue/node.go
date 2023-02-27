package queue

import "edd.com/samples/structs"

type NodeQueue struct {
	Student *structs.Student
	Next    *NodeQueue
}

func (n *NodeQueue) PrintNode() {
	n.Student.Print()
}
