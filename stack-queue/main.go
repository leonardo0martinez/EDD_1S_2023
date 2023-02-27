package main

import (
	"fmt"

	"edd.com/samples/funcs"
	"edd.com/samples/queue"
	"edd.com/samples/stack"
)

func main() {

	// LOGS ADMIN

	//Pila de Espera
	espera := &stack.Stack{Size: 1} //PARA EL PROYECTO: COLA
	// Usuario Registrados
	registrados := &queue.Queue{Size: 1} //PARA EL PROYECTO: LISTA DOBLE

	var (
		user     string
		password string
		exit     bool
	)

	option := 0

	for !exit {
		fmt.Println("Ejemplo para Proyecto ")
		fmt.Println("______________________________")
		fmt.Println("1. Iniciar Sesion")
		fmt.Println("2. Salir del Sistema")
		fmt.Println("______________________________")
		fmt.Print("Elige una opcion: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Print("Ingresa tu usuario: ") //NUM DE CARNET
			fmt.Scanln(&user)
			fmt.Print("Ingresa tu Password: ")
			fmt.Scanln(&password)
			if user == "admin" && password == "admin" {
				funcs.RegistrarEstudiante(espera, registrados)
			} else {
				fmt.Print("Buscar Aceptado")
			}
		case 2:
			fmt.Print("Adios !")
			exit = true
		}

	}

}
