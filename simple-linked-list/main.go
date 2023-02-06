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
	fmt.Println("Eliminando el primero:")
	list.Delete(1)
	list.Print()
	fmt.Println("Eliminando el último:")
	list.Delete(10)
	list.Print()
	fmt.Println("Eliminando entre nodos:")
	list.Delete(5)
	list.Print()

}
