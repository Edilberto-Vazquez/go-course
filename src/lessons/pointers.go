package lessons

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() { // metodo paso por referencia
	myPC.ram = myPC.ram * 2
}

func duplicateRAMFunc(myPC *pc) { // funcion paso por referencia
	myPC.ram = myPC.ram * 2
}

func Pointers() {
	a := 50
	b := &a // accede a la direccion de memoria

	fmt.Println(b)
	fmt.Println(*b) // accede al valor almacenado en memoria

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "MSI"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
	duplicateRAMFunc(&myPC)

	fmt.Println(myPC)
}
