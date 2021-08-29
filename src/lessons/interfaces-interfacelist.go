package lessons

import "fmt"

type shapes2D interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base   float64
	altura float64
}

func (s square) area() float64 {
	return s.base * s.base
}

func (r rectangle) area() float64 {
	return r.base * r.altura
}

func calculate(f shapes2D) {
	fmt.Println("Area:", f.area())
}

func InterfacesInterfacesList() {
	mySquare := square{base: 2}
	myRectangle := rectangle{base: 2, altura: 4}

	calculate(mySquare)
	calculate(myRectangle)

	// Lista de interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)
}
