package funcs

import (
	"fmt"

	"edd.com/samples/queue"
	"edd.com/samples/stack"
)

func RegistrarEstudiante(stack *stack.Stack, queue *queue.Queue) {
	option := 0
	validar := 0
	var (
		carnet   int
		name     string
		password string
	)
	for option != 4 {
		fmt.Println("Ejemplo para Proyecto ")
		fmt.Println("______________________________")
		fmt.Println("1. Registrar Estudiante")
		fmt.Println("2. Aceptar o Rechazar Estudiantes")
		fmt.Println("3. Estudiantes Aceptados")
		fmt.Println("4. Cerrar Sesión")
		fmt.Println("______________________________")
		fmt.Print("Elige una opcion: ")
		fmt.Scanln(&option)
		switch option {
		//---------------------------------------------------------------------
		// REGISTRAR ESTUDIANTE------------------------------------------------
		//---------------------------------------------------------------------
		case 1:
			fmt.Println("______________________________")
			fmt.Println("Regitrar Estudiante")
			fmt.Println("______________________________")
			fmt.Print("Ingresa tu Carnet: ")
			fmt.Scanln(&carnet)
			fmt.Print("Ingresa tu Nombre: ")
			fmt.Scanln(&name)
			fmt.Print("Ingresa tu Password: ")
			fmt.Scanln(&password)
			stack.Push(carnet, name, password)
			fmt.Print("Agregado a la Pila de Espera.")
			fmt.Println("______________________________")
		//---------------------------------------------------------------------
		// ACEPTAR ESTUDIANTE------------------------------------------------
		//---------------------------------------------------------------------
		case 2:
			fmt.Println("______________________________")
			fmt.Println("Aceptar Estudiante")
			fmt.Println("______________________________")
			for validar != 3 {
				if !stack.IsEmpty() {
					temp, size := stack.Pop()
					fmt.Printf("___________________Pendientes : %d\n", size)
					fmt.Printf("Carnet: %d\n", temp.Carnet)
					fmt.Printf("Nombre: %s\n", temp.Name)
					fmt.Println("1. Aceptar")
					fmt.Println("2. Rechazar")
					fmt.Println("3. Regresar")
					fmt.Println("______________________________")
					fmt.Print("Elige una opcion: ")
					fmt.Scanln(&validar)
					switch validar {
					case 1:
						// Agregar a la cola
						queue.Enqueue(temp)
						fmt.Println("Estudiante Aceptado")
					case 2:
						fmt.Println("Estudiante Rechazado")
					}
				} else {
					validar = 3
					fmt.Println("No hay Estudiantes en epera")
				}
			}
		case 3:
			//---------------------------------------------------------------------
			// ESTUDIANTES ACEPTADOS------------------------------------------------
			//---------------------------------------------------------------------
			fmt.Println("______________________________")
			fmt.Println("Estudiantes Aceptados")
			fmt.Println("______________________________")
			queue.Print()
		}
	}
}
