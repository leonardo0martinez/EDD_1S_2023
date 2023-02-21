package main

import (
	"encoding/csv" // Librería para la lectura de CSV
	"fmt"
	"log"
	"os"
)

// Recibe la ruta del archivo como parámetro
func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Archivo no existe "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Error en la lectura del archivo"+filePath, err)
	}
	return records
}

func main() {
	records := readCsvFile("./prueba.csv")

	// Descripción del array
	// [[carnet, nombre, contraseña], [201780044,Leonardo Martinez,leo1234], ....]
	for index, row := range records { // EJEMPLO DE FOR EACH EN GOLANG
		fmt.Printf("%d: %s\t%s\n", index, row[0], row[1])
	}
}
