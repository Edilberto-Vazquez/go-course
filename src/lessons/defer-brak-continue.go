package lessons

import "fmt"

func DeferBreakContinue() {
	// Defer
	defer fmt.Println("Mundo")
	fmt.Println("Hola")

	// continue and break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// break
		if i == 2 {
			fmt.Println("Es 5")
			break
		}
	}
}
