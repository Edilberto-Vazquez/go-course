package lessons

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done() // se manda la señal de que se termino la operacion
	fmt.Println(text)
}

func Goroutines() {

	var wg sync.WaitGroup // se crea un waitGroup

	fmt.Println("Hello")
	wg.Add(1) // se agrega la go routine al wait group
	go say("World", &wg)
	wg.Wait() // El programa espera a que terminen las goroutines

	go func(text string) {
		fmt.Println("Adios")
	}("Adios")

	time.Sleep(time.Second * 1) // mala practica usar sleep

}
