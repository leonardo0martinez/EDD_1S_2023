package circular

import (
	"fmt"
	"strconv"
)

type CircularList struct {
	head *Node
	tail *Node
}

// INSERTAR AL FINAL
func (list *CircularList) AddEnd(value int) {

	newNode := &Node{value: value, next: nil}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		// Asignar puntero a la cabeza haciendo la lista circular
		newNode.next = list.head.next
	} else {
		list.tail.next = newNode
		list.tail = newNode
		list.tail.next = list.head
	}
}

// INSERTAR AL FINAL
func (list *CircularList) AddFront(value int) {
	newNode := &Node{value: value, next: nil}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		// Asignar puntero a la cabeza haciendo la lista circular
		newNode.next = list.head.next
	} else {
		newNode.next = list.head
		list.head = newNode
		list.tail.next = newNode
	}
}

// MÉTODO PARA IMPRIMIR LA LISTA
func (list *CircularList) Print() {
	temp := list.head
	for temp.next != list.head {
		fmt.Printf("%d, ", temp.value)
		temp = temp.next
	}
	fmt.Printf("%d", temp.value)
}

// MÉTODO PARA GENERAR CÓDIGO GRPHVIZ
func (list *CircularList) Graph() {
	temp := list.head
	conn := ""
	nodes := ""
	counter := 0
	for temp.next != list.head {
		nodes += "N" + strconv.Itoa(counter) + "[label=\"Valor:" + strconv.Itoa(temp.value) + "\"];\n"
		conn += "N" + strconv.Itoa(counter) + "->"
		temp = temp.next
		counter++
	}
	nodes += "N" + strconv.Itoa(counter) + "[label=\"Valor:" + strconv.Itoa(temp.value) + "\"];\n"
	conn += "N" + strconv.Itoa(counter) + "-> N0"
	fmt.Println(nodes)
	fmt.Println(conn)
}
