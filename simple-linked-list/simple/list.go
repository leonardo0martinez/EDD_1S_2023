package simple

import (
	"fmt"
)

// Declaración de la estructura
type SimpleList struct {
	head *Node
}

// Se agrega el puntero hacia el struct para hacerlo parte de el
func (list *SimpleList) Insert(value int) {
	//Declarar nuevo nodo
	newNode := &Node{value: value, next: nil}
	//Verificar si la lista está vacía
	if list.head == nil {
		list.head = newNode
	} else {
		//Recorrer hasta encontrar el último nodo
		temp := list.head
		for temp.next != nil {
			temp = temp.next
		}
		//Agregar el nuevo nodo hasta el final
		temp.next = newNode
	}
}

// Método para imprimir la lista
func (list SimpleList) Print() {
	temp := list.head
	for temp.next != nil {
		fmt.Printf("%d, ", temp.value)
		temp = temp.next
	}
	fmt.Printf("%d\n", temp.value)
}

// Método para buscar un valor en la lista
func (list *SimpleList) FindIndex(value int) (result bool, index int) {
	temp := list.head
	index = 0
	result = false
	for temp != nil {
		if temp.value == value {
			result = true
		}
		index++
		temp = temp.next
	}

	if !result {
		index = -1
	}
	return result, index
}

// Método para eliminar valor de la lista
func (list *SimpleList) Delete(value int) (result bool) {
	result = false
	// Si es el primero de la lista
	if list.head.value == value {
		result = true
		if list.head.next == nil {
			list.head = nil
		} else {
			list.head = list.head.next
		}
	} else {
		temp := list.head
		// Dos casos para eliminar [Se eliminará el 2]
		// 1 -> |2| -> 3 		----- Que el siguiente nodo no sea nulo
		// 1 -> |2| -> Null		----- Que el siguiente nodo sea nulo
		for temp != nil {
			//Encontrar el nodo anterior al que se elimina
			if temp.next != nil {
				if temp.next.value == value {
					break
				}
			}
			temp = temp.next
		}
		if temp == nil {
			return false
		} else {
			//Se asigna el apuntador al siguiente del que se elimina
			// 1 -> |2| -> 3   		==>  1 -> 3
			// 1 -> |2| -> Null   	==>  1 -> Null (toma el valor del apuntando a nulo)
			temp.next = temp.next.next
			result = true
		}
	}
	return result
}

func (list *SimpleList) Update(value int, newValue int) {
	temp := list.head
	for temp != nil {
		if temp.value == value {
			break
		}
		temp = temp.next
	}
	temp.value = newValue
}

func (list *SimpleList) GetNode(value int) *Node {
	temp := list.head
	for temp != nil {
		if temp.value == value {
			break
		}
		temp = temp.next
	}
	return temp
}
