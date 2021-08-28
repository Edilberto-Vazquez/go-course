package lessons

import "fmt"

func AritmeticOperators() {
	// area del cuadrado
	const baseCudrado = 10
	areacuadrado := baseCudrado * baseCudrado
	fmt.Println("Area cuadrado:", areacuadrado)

	// suma
	x := 10
	y := 12
	result := x + y
	fmt.Println("suma: ", result)

	// resta
	result = x - y
	fmt.Println("resta: ", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion: ", result)

	// division
	result = y / x
	fmt.Println("division: ", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	// incremental
	x++
	fmt.Println("Incremental: ", x)

	// decremental
	x--
	fmt.Println("Decremental: ", x)
}
