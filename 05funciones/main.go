package main

import "fmt"

func main() {
	nombre := "Vincent"

	// Funciones
	saludar(nombre)
	fmt.Println()
	multiplica(2, 3, "Resultado")
	fmt.Println()
	result := suma(1, 3)
	fmt.Println(result)
	fmt.Println()

	//Recibir double return
	value1, value2 := doubleReturn(4)
	fmt.Println(value1, value2)

	// Descartar un valor
	value3, _ := doubleReturn(4)
	fmt.Println(value3)

}

// Declaracion
func saludar(nombre string) {
	fmt.Printf("Hola %s", nombre)
}

// Funcion con tipos repetidos
func multiplica(valor, cantidad int32, texto string) {
	fmt.Println(texto, valor*cantidad)
}

// Retorna un valor
func suma(a int32, b int32) int32 {
	return a + b
}

// Podemos retornar mas de un valor
func doubleReturn(a int) (c, d int) {
	return a, a * 2
}
