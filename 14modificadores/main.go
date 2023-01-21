package main

import "fmt"
// Se puede agregar un alias para poder acceder mas facil
import ( pk "this-is-go/14modificadores/mypackete" )

func main() {
	polera := pk.Polera
	polera.Talla = "L"

	result := pk.Sumar(2, 4)

}
