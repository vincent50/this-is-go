package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	// If - Else
	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// If And
	if valor1 == 1 && valor2 == 3 {
		fmt.Println("Es verdad")
	}

	// If OR
	if valor1 == 0 || valor2 == 3 {
		fmt.Println("Es verdad")
	}

	// Condicional nil
	value, err := strconv.Atoi("b3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value", value)

}
