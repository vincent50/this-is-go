package main

import "fmt"

type PC struct {
	ram   int
	disk  int
	brand string
}

// Podemos sobreescribir el metodo to string como lo hacemos en otros lenguajes
// en este caso debemos definir solamente el metodo ya existente String el cual se utiliza para imprimir su informacion
func (pc PC) String() string {
	return fmt.Sprintf("Mi PC tiene %d de ram, %d de disco y es de la marca %s", pc.ram, pc.disk, pc.brand)
}

// este mismo metodo sirve para utilizar en otros tipos de metodos

func main() {
	myPC := PC{ram: 64, disk: 1240, brand: "HP"}
	fmt.Println(myPC)
}
