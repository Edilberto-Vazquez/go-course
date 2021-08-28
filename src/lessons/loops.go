package lessons

import "fmt"

func Loops() {
	// For condicional
	for i := 0; i < 10; i++ {
		fmt.Println((i))
	}

	fmt.Printf("\n")

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// }

	// For range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
