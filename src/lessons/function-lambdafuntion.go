package lessons

import "fmt"

func NormalFunction(message string) {
	fmt.Println(message)
}

func TripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func ReturnValue(a int) int {
	return a * 2
}

func DoubleReturn(a int) (c, d int) {
	return a, a * 2
}
