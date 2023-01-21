package main

func main() {
	// Diccionarion
	personas := make(map[string]int)
	personas["Vincet"] = 32
	personas["Ciri"] = 3
	personas["Belen"] = 29

	// Recorrer map
	for index, value := range personas {
		fmt.Println("Index:",index)
		fmt.Println("Valor:",value)
	}

	// Encontrar un valor
	value := personas["Ciri"]
	fmt.Println("Valor:",value)

	// Al no existir un valor Go retornara el Zero value del tipo de dato
	// En este caso puede devolver un 0
	// para controlarlo, podemos usar el segundo valor retornado el cual es un booleano
	nombre, ok := personas["falso"]
	if ok {
		fmt.Println("Valor:",nombre)
	}
}
