package main

import "fmt"

// defer finaliza al terminar el bloque de codigo en donde existe
func deferTest() {
	defer fmt.Println("Final de la funcion")
}

func main() {
	// defer
	// Ejecuta defer al finalizar el proceso
	defer fmt.Println("Hola")
	fmt.Println("Mundo")
	deferTest()

	// Continue
	// Permite continuar con la siguiente iteracion sin ejecutar el codigo siguiente
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Break
	// Detiene el bucle al llegar al keyword break
	contador := 1
	for {
		fmt.Println(contador)
		if contador == 10 {
			break
		}
		contador++
	}
}
