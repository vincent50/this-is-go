package main

import (
	"fmt"
	"sync"
)

// Para agregar una funcion en una goroutine se especifica como tipo de dato
// para al finalizar enviar una señal de que esta funcion a finalizado
func say(text string, wg *sync.WaitGroup) {
	fmt.Println(text)
	wg.Done()
}

func main() {
	// Con esto declaramos un grupo de goroutines
	wg := sync.WaitGroup{}

	fmt.Println("Hello")
	wg.Add(1) // Aqui especificamos cuantas tenemos para esperar

	// Ejecutamos la goroutine
	go say("world", &wg)

	wg.Wait() // Esperamos a que termine la goroutine

	// Otra forma es utilizando funciones anonimas
	go func(text string) {
		fmt.Println(text)
		wg.Done()
	}("Adios")

	wg.Add(1)
	wg.Wait()
}
