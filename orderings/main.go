package main

import (
	"fmt"
)

func main() {

	var disorder = []int{10, 12, 10, 212, 3, 14, 44, 30}
	fmt.Println("Array desordenado: ")
	fmt.Println(disorder)
	fmt.Println("Ordenamiento Burbuja: ")
	fmt.Println(bubbleSort(disorder))
	fmt.Println("Ordenamiento por Inserción: ")
	fmt.Println(insertionSort(disorder))
	fmt.Println("Ordenamiento por Quicksort:")
	final := len(disorder) - 1
	fmt.Println(quickSort(disorder, 0, final))

}

// ORDENAMIENTO BURBUJA
func bubbleSort(list []int) []int {
	var temp int
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			// COMPARA TODOS LOS VALORES DE LA LISTA
			// fmt.Printf("%d > %d\n", list[i], list[j])
			if list[i] < list[j] {
				temp = list[i]
				list[i] = list[j]
				list[j] = temp
			}
		}
	}
	return list
}

// ORDENAMIENTO POR INSERCIÓN
func insertionSort(list []int) []int {
	var temp int
	for i := 1; i < len(list); i++ {
		temp = list[i]
		for j := i - 1; j >= 0 && list[j] > temp; j-- {
			list[j+1] = list[j]
			list[j] = temp
		}
	}
	return list
}

// ORDENAMIENTO POR INSERCIÓN
func quickSort(list []int, left int, right int) []int { // Left y Right son posiciones del array
	pivote := list[left] // TOMAR EL PRIMER ELEMENTO COMO PIVOTE
	i := left            // i REALIZA LA BÚSQUEDA DE IZQUIERDA A DERECHA
	j := right           // j REALIZA LA BÚSQUEDA DE DERECHA A IZQUIERDA
	var temp int         // GUARDAR VALOR PARA COMPARAR

	for i < j { // MINTRAS NO SE CRUCEN LAS BÚSQUEDAS

		for list[i] <= pivote && i < j { // BUSCABA EL ELEMENTO MAYOR AL PIVOTE
			i++
		}

		for list[j] > pivote { // BUSCABA EL ELEMENTO MENOR AL PIVOTE
			j--
		}

		if i < j {
			temp = list[i]
			list[i] = list[j]
			list[j] = temp
		}
	}

	list[left] = list[j] // se coloca el pivote en su lugar de forma que tendremos
	list[j] = pivote     // los menores a su izquierda y los mayores a su derecha
	if left < j-1 {
		quickSort(list, left, j-1) // ordenamos subarray izquierdo
	}

	if j+1 < right {
		quickSort(list, j+1, right) // ordenamos subarray derecho
	}

	return list
}
