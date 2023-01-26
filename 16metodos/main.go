package main

import "fmt"

type PC struct {
	ram   int
	disk  int
	brand string
}

// Para agregar un metodo a un struct se debe realizar externo a la estructura
// pero especificamos entre parentesis antes del nombre de la funcion
func (pc PC) ping() {
	fmt.Println("PONG")
}

// Para poder modificar valores podemos especificar el espacio de memoria
func (pc *PC) duplicateRam() {
	pc.ram *= 2
}

func main() {
	myPC := PC{ram: 128}
	myPC.ping() // Luego podemos llamarlo como lo usamos en otros lenguajes

	fmt.Println(myPC.ram)
	myPC.duplicateRam()
	fmt.Println(myPC.ram)

}
