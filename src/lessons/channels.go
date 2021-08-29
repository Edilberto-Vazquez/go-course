package lessons

import (
	"fmt"
)

func sayC(text string, c chan<- string) {
	c <- text
}

func Channels() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go sayC("Bye", c)

	fmt.Println(<-c)
}
