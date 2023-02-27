package queue

import "edd.com/samples/structs"

type Queue struct {
	head *NodeQueue
	Size int
}

// AGREGAR A LA LISTA AL FINAL
func (queue *Queue) Enqueue(std *structs.Student) {
	//Declarar nuevo nodo
	newNode := &NodeQueue{Student: std, Next: nil}
	//Verificar si la lista está vacía
	if queue.head == nil {
		queue.head = newNode
	} else {
		//Recorrer hasta encontrar el último nodo
		temp := queue.head
		for temp.Next != nil {
			temp = temp.Next
		}
		//Agregar el nuevo nodo hasta el final
		temp.Next = newNode
		queue.Size += 1
	}
}

// ELIMINAR AL PRINCIPIO
func (queue *Queue) Dequeue() (*structs.Student, int) {
	if queue.head == nil {
		return nil, -1
	} else {
		// Tomar valor de la
		temp := queue.head
		// Eliminar de la Pila
		queue.head = queue.head.Next
		queue.Size -= 1
		// Retornar Estudiante
		return temp.Student, queue.Size
	}
}

// METODO PARA IMPRIMIR LA COLA
func (queue *Queue) Print() {
	temp := queue.head
	for temp != nil {
		temp.Student.Print()
		temp = temp.Next
	}
}
