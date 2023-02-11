package doubly

type Node struct {
	Carnet int
	Nombre string
	Next   *Node
	Prev   *Node
}

func (n *Node) GetCarnet() int {
	return n.Carnet
}

func (n *Node) GetNombre() string {
	return n.Nombre
}

func (n *Node) SetNombre(str string) {
	n.Nombre = str
}

func (n *Node) SetCarnet(carnet int) {
	n.Carnet = carnet
}
