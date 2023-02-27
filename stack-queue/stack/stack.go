package stack

import "edd.com/samples/structs"

type Stack struct {
	head *NodeStack
	Size int
}

func (stack *Stack) Push(carnet int, name string, password string) {
	// Crean el estudiante
	newStd := &structs.Student{Carnet: carnet, Name: name, Password: password}
	// Crear el nuevo nodo
	newNode := &NodeStack{Student: newStd, Next: nil}

	if stack.head == nil {
		stack.head = newNode
	} else {
		// Inserción al inicio
		temp := stack.head
		stack.head = newNode
		newNode.Next = temp
		stack.Size += 1
	}
}

func (stack *Stack) Pop() (*structs.Student, int) {
	if stack.head == nil {
		return nil, -1
	} else {
		// Tomar valor de la
		temp := stack.head
		// Eliminar de la Pila
		stack.head = stack.head.Next
		stack.Size -= 1
		// Retornar Estudiante
		return temp.Student, stack.Size
	}
}

func (stack *Stack) Print() {
	temp := stack.head
	for temp != nil {
		temp.Student.Print()
		temp = temp.Next
	}
}

func (stack *Stack) IsEmpty() bool {
	return stack.head == nil
}
