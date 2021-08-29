package lessons

import (
	"fmt"
	"strings"
)

func LoopThroughSlice() {
	slcie := []string{"Hola", "que", "hace"}

	for i, v := range slcie {
		fmt.Println(i, v)
	}
}

func IsPalindromo(text string) {
	var textReverse string
	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}
