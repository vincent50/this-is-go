package main

import "fmt"

func main() {
	fmt.Println("Variables!")
	// Puede cambiar su valor
	base := 12          // Declaracion de variables enteras
	var altura int = 14 // Declaracion estandar
	var area int        // Declarada pero no inicializada
	fmt.Println("Variables:", base, altura, area)
	fmt.Println()

	fmt.Println("Constantes!")
	// No puede cambiar su valor
	const PI float64 = 3.14 // Tipeado
	const PI2 = 3.14        // No tipeado
	fmt.Println("Constantes:", PI, PI2)
	fmt.Println()

	fmt.Println("Zero values!")
	// Go agrega un valor por default en caso de no inicializar
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("Defaults:", a, b, c, d)
	fmt.Println()

}
