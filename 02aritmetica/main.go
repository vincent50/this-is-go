package main

import "fmt"

func main() {
	x := 10
	y := 50

	// SUMA
	result := x + y
	fmt.Println("SUMA Resultado:", result)

	// RESTA
	result = y - x
	fmt.Println("RESTA Resultado:", result)

	// MULTIPLICACION
	result = x * y
	fmt.Println("MULTIPLICACION Resultado:", result)

	// DIVISION
	result = y / x
	fmt.Println("DIVISION Resultado:", result)

	// MOD
	result = y % x
	fmt.Println("MOD Resultado:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado:", areaCuadrado)

}
