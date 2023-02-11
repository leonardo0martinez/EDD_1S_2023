package main

import (
	"edd.com/samples/doubly"
)

func main() {

	list := &doubly.DoublyList{}

	list.InsertEnd(201780044, "Leonardo1 Martinez")
	list.InsertEnd(201780044, "Leonardo2 Martinez")
	list.InsertEnd(201780044, "Leonardo3 Martinez")
	list.InsertEnd(201780044, "Leonardo4 Martinez")

	list.Print()
	list.Graph()
	//list.Print2()

}
