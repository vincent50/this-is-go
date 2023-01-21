package main

type Car struct {
	brand string
	year int
}

func main() {
	// Forma de crear clases en Go
	myCar := Car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	otherCar := Car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
	// para los campos faltantes se agregan los Zero values
}
