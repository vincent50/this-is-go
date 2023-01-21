package mypackage

// Struct publica: se debe crear con la primera palabra en mayuscula y definir un comentario justo arriba

// Polera objeto para referirse a una polera
type Polera struct {
	Talla string // Field publico, inicia con mayuscula
	color string // Field privado, inicia con minuscula
}

// Struct privada: se debe crear con la primera palabra en minuscula y no es necesario agregar un comentario

type car struct {
	brand string
	year int
}

// Este metodo tambien se puede utilizar con funciones

// Sumar funcion que permite sumar dos numeros
func Sumar(a,b int) int {
	return a + b
}

func sumaSecreta(a,b int) int {
	return a + b
}

