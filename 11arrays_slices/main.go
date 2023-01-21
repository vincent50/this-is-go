package main

import "fmt"

func main() {
	// Array
	// tiene valor fijo, inmutable
	var array [4]int

	fmt.Println(array)
	// pero podemos modificar sus valores
	array[0] = 1
	array[1] = 2
	fmt.Println(array)

	// Manejo
	// len: cantidad del array
	// cap: capacidad del array
	fmt.Println(array, len(array), cap(array))

	// SLICE
	// no tiene valor fijo, es mutable
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos en el slice
	// El ultimo elemento es exclusivo
	// El primer elemento es inclusivo
	fmt.Println(slice[0])   // un elemento
	fmt.Println(slice[:3])  // hasta el indice 2
	fmt.Println(slice[2:4]) // entre el indice 2 y 3
	fmt.Println(slice[4:])  // desde el indice 4

	// Append
	// Agregar elementos en un slice
	slice = append(slice, 8)
	fmt.Println(slice, len(slice), cap(slice))

	// Agregar una lista
	newSlice := []int{9, 10, 11}
	// ... similar a javascript, lo desectructura y lo agrega
	slice = append(slice, newSlice...)
	fmt.Println(slice, len(slice), cap(slice))

	// Recorrer un Slice
	slice := []string{"manzana", "uva", "naranja", "platano"}

	// range devuelve dos valores, el indice y el valor de cada iteracion
	for index, value := range slice {
		fmt.Println("Index:",index)
		fmt.Println("Valor:",value)
	}

	// Si solo queremos rescatar el valor podemos agregar un _ para no tomar el index
	for _, value := range slice {
		fmt.Println("Valor:",value)
	}

	// si solo necesitamos solo en indice no necesitamos agregar _ ya que index es el valor por defecto
	for index := range slice {
		fmt.Println("Index:",index)
	}

}
