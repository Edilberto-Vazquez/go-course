package lessons

import "fmt"

func ConstVarsZerovalues() {
	// declaracion de constantes
	fmt.Println(("Declaracion de constantes"))
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi1:", pi)
	fmt.Println("pi2:", pi2)

	// declaracion de variables enteras
	fmt.Println("declaracion de variables enteras")
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// zero values
	fmt.Println("zero values")
	var entero int
	var decimal float64
	var cadena string
	var boleano bool

	fmt.Println(entero, decimal, cadena, boleano)
}
