package dot

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// MÉTODO PARA ESCRIBIR EL ARCHIVO ".dot"
func WriteDotFile(code string, fileName string, path string) {
	// Verifica que el archivo existe
	var _, err = os.Stat(path + "\\" + fileName)
	// Crea el archivo si no existe
	if os.IsNotExist(err) {
		// Si no existe lo crea
		var file, err = os.Create(fileName)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer file.Close()
	} else {
		// Si existe lo elimina, para crearlo de nuevo
		// Y actualizar el archivo si fuese necesario
		err := os.Remove(fileName)
		if err == nil {
			var file, err = os.Create(fileName)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer file.Close()
		}
	}

	// Abre archivo usando permisos de escritura
	var file, _ = os.OpenFile(fileName, os.O_RDWR, 0644)
	_, err = file.WriteString(code)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Guardar los cambios
	err = file.Sync()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Archivo .dot creado.")

}

// Método para ejecutar comando en consola
func GeneratePNG(fileName string, path string) {
	// COMANDO A EJECUTAR: dot -Tpng {RutaArchivo.dot} -o {RutaImagen.png}
	// Ejemplo: dot -Tpng archivito.dot -o prueba.png 				#Archivo en el mismo directorio
	// _, err := exec.Command(path.Join(path.Dir(currPath), fileName), "-Tpng", fileName).Output()
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fixedPath := strings.ReplaceAll(path)
	// cmd := exec.Command("dot -Tpng " + fileName + " -o "  + strings.Replace(fileName, ".dot", ".png", -1))
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	path2, _ := exec.LookPath("dot")
	cmd, err := exec.Command(path2, "dot", "-Tpng", fileName).Output()
	if err != nil {
		fmt.Print(err)
	}
	mode := int(0777)
	os.WriteFile(strings.Replace(fileName, ".dot", ".png", -1), cmd, os.FileMode(mode))
}
