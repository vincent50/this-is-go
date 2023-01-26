package main

// Se puede agregar un alias para poder acceder mas facil
// en este caso pk es el nombre del alias
import (
	"fmt"
	pk "this-is-go/14modificadores/mypackage"
)

func main() {
	polera := pk.Polera{}
	polera.Talla = "L"

	result := pk.Sumar(2, 4)
	fmt.Println(result)
}
