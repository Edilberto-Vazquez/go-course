package lessons

import "fmt"

func ConditionalIf() {
	valor1 := 2
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// With and
	if valor1 == 2 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	// with or
	if valor1 == 1 || valor2 == 2 {
		fmt.Println("Es verdad")
	}

	// with not
	if valor1 == 2 && valor2 == 2 {
		fmt.Println("Es verdad")
	}
}
