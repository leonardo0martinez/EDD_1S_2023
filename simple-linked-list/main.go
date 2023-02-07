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
	fmt.Printf("Eliminando el primero: %t\n", list.Delete(1))
	list.Print()
	fmt.Printf("Eliminando el último: %t\n", list.Delete(10))
	list.Print()
	fmt.Printf("Eliminando entre nodos: %t\n", list.Delete(5))
	list.Print()
	fmt.Printf("Eliminando uno que no exite: %t\n", list.Delete(1000))
	list.Print()

}
