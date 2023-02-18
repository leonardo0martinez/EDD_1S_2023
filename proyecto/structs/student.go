package structs

import "fmt"

type Student struct {
	Carnet   int
	Name     string
	Password string
}

func (s *Student) Print() {
	fmt.Println("________________________________________________")
	fmt.Printf("Carnet: %d\n", s.Carnet)
	fmt.Printf("Nombre: %s\n", s.Name)
}
