package doubly

import (
	"fmt"
	"strconv"
)

type DoublyList struct {
	head *Node
	tail *Node
}

// INSERTAR AL AL FINAL
func (list *DoublyList) InsertEnd(carnet int, nombre string) {
	newNode := &Node{Carnet: carnet, Nombre: nombre, Prev: nil, Next: nil}
	// Verificar que la lista está vacía
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		// Apuntador del nodo nuevo al final
		list.tail.Next = newNode
		// Apuntador del nodo nuevo a la cola
		newNode.Prev = list.tail
		// Cabiar la cola al nodo nuevo
		list.tail = newNode
	}
}

// INSERTAR AL INICIO
func (list *DoublyList) InsertStart(carnet int, nombre string) {
	newNode := &Node{Carnet: carnet, Nombre: nombre, Prev: nil, Next: nil}
	// Verificar que la lista está vacía
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		// Apuntador del nodo nuevo al inicio
		list.head.Prev = newNode
		// Apuntador del nodo nuevo a la cabeza
		newNode.Next = list.head
		// Cabiar la cabeza al nodo nuevo
		list.head = newNode
	}
}

// IMPRIMIR
func (list *DoublyList) Print() {
	temp := list.head
	for temp.Next != nil {
		fmt.Printf("Nombre: %s, \n", temp.GetNombre())
		temp = temp.Next
	}
	fmt.Printf("Nombre: %s, \n", temp.GetNombre())
}

// IMPRIMIR AL REVÉS
func (list *DoublyList) Print2() {
	temp := list.tail
	for temp.Prev != nil {
		fmt.Printf("Nombre: %s, \n", temp.GetNombre())
		temp = temp.Prev
	}
	fmt.Printf("Nombre: %s, \n", temp.GetNombre())
}

// GENERAR CÓDIGO GRAPHVIZ
func (list *DoublyList) GraphCode() string {
	temp := list.head
	nodes := ""
	conn := ""
	counter := 0
	for temp.Next != nil {
		nodes += "N" + strconv.Itoa(counter) + "[label=\"Carnet:" + strconv.Itoa(temp.Carnet) + "\nNombre: " + temp.Nombre + "\"];\n"
		conn += "N" + strconv.Itoa(counter) + "->"
		temp = temp.Next
		counter++
	}
	nodes += "N" + strconv.Itoa(counter) + "[label=\"Carnet:" + strconv.Itoa(temp.Carnet) + "\nNombre: " + temp.Nombre + "\"];\n"
	conn += "N" + strconv.Itoa(counter) + "\n"
	temp = list.tail
	for temp.Prev != nil {
		conn += "N" + strconv.Itoa(counter) + "->"
		temp = temp.Prev
		counter--
	}
	conn += "N" + strconv.Itoa(counter)

	return "digraph G {\n" +
		"node[shape=rectangle, style=filled, color=lightsalmon];\n" +
		"rankdir=LR;\n" +
		nodes + // NODOS
		conn + // CONEXIONES
		"\n}"
}
