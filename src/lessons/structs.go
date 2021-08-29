package lessons

import "fmt"

type car struct {
	brand string
	year  int
}

func Structs() {
	// First form
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Second form
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
