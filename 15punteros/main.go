package main

import "fmt"

func main() {
	a := 50
	b := &a // para obtener el espacio de memoria utilizamos &

	fmt.Println(b)
	fmt.Println(*b)

	// Para referenciar el espacio de memoria utilizamos *
	*b = 100
	fmt.Println(a)

}
