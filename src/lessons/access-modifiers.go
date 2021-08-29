package lessons

import "fmt"

// CarPublic Car with public access
type CarPublic struct {
	Brand string
	Year  int
}

// carPrivate Car with private access

type carProvate struct {
	brand string
	year  int
}

// PrintMessage print a message
func PrintMessage() {
	fmt.Println("Hola")
}
