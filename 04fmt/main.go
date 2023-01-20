package main

import "fmt"

func main() {
	// Uso de package fmt
	helloMessage := "Hello"
	worldMessage := "world"

	// Printl imprimir con salto de linea
	fmt.Println(helloMessage, worldMessage)

	// Printf String format
	nombre := "Pokemones"
	cantidad := 300
	fmt.Printf("existen mas de %d %s en la actualidad", cantidad, nombre)
	// Tambien se puede imprimir sin saber el tipo de dato, pero no es buena practica
	fmt.Printf("existen mas de %v %v en la actualidad", cantidad, nombre)

	// Sprintf guardar el nuevo string en una variable
	message := fmt.Sprintf("existen mas de %d %s en la actualidad", cantidad, nombre)
	fmt.Println(message)

	// Saber el tipo de dato
	fmt.Printf("message: %T", message)
	fmt.Println()

	fmt.Printf("cantidad: %T", cantidad)
	fmt.Println()

}
