package main

import (
	"log"
	"os"

	"edd.com/samples/dot"
	"edd.com/samples/doubly"
)

func main() {

	list := &doubly.DoublyList{}

	list.InsertEnd(201700000, "Madara Uchiha")
	list.InsertEnd(201700001, "Obito Uchiha")
	list.InsertEnd(201700002, "Kabuto Yakushi")
	list.InsertEnd(201700003, "Papeles Konan")
	list.InsertEnd(201700004, "Judío Hidan")
	list.InsertEnd(201700005, "Orochimaru Snake")
	list.InsertEnd(201700006, "Marionette Sasori")
	list.InsertEnd(201700007, "Kakuzu Mooney")
	list.InsertEnd(201700008, "Onion Nagato")
	list.InsertEnd(201700009, "Kisame Hoshigaki")
	list.InsertEnd(201700010, "Deidara Pum")
	list.InsertEnd(201700011, "Itachi Uchiha")
	list.InsertEnd(201700012, "Aloevera Zetsu")

	// Imprimir código graphviz
	// fmt.Print(list.GraphCode())

	// Obtener la ruta de directorio actual para que no de clavos xd
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	// Escribir el archivo .dot
	dot.WriteDotFile(list.GraphCode(), "prueba.dot", path)
	// Ejecutar COmando en consola
	dot.GeneratePNG("prueba.dot", path)

}
