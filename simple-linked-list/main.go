package main

import (
	"fmt"

	"edd.com/samples/simple"
)

func main() {

	//Declarar Lista Enlazada
	list := simple.SimpleList{}

	fmt.Println("Lista :")
	for i := 1; i <= 10; i++ {
		list.Insert(i)
	}
	list.Print()

	// Segundo Método
	fmt.Println("Actualizando valor desde la clase Lista: ")
	list.Update(10, 20)
	list.Print()

	fmt.Println("Actualizando valor desde la clase Nodo: ")
	temp := list.GetNode(4)
	temp.SetValue(50)
	list.Print()
	fmt.Println("Imprimiendo el valor del Nodo: ")
	fmt.Println(temp.GetValue())
	// fmt.Print()
	// fmt.Printf("Eliminando el primero: %t\n", list.Delete(1))
	// list.Print()
	// fmt.Printf("Eliminando el último: %t\n", list.Delete(10))
	// list.Print()
	// fmt.Printf("Eliminando entre nodos: %t\n", list.Delete(5))
	// list.Print()
	// fmt.Printf("Eliminando uno que no exite: %t\n", list.Delete(1000))
	// list.Print()

}
