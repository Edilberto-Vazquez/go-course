package lessons

import "fmt"

func KeyValue() {
	m := make(map[string]int)

	m["jose"] = 14
	m["pepito"] = 23

	fmt.Println(m)

	// recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar valor
	value, ok := m["jose"]
	fmt.Println(value, ok)
}
