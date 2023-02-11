package main

import "edd.com/samples/circular"

func main() {

	list := circular.CircularList{}

	for i := 1; i <= 10; i++ {
		list.AddEnd(i)
	}
	list.Print()

	list.Graph()

}
