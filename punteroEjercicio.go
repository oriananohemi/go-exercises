package punteroEjercicio

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func puntero() {
	a := 50
	// el & es para acceder al lugar en memoria
	b := &a

	fmt.Println(b)
	// el * es para acceder al valor en memoria
	fmt.Println(*b)

	// Si modificamos una variable con *, todas las variables que esten aputando a este lugar cambiaran
	*b = 100
}

// Actualiza la ram
func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func main() {
	myPC := pc{ram: 16, disk: 200, brand: "msi"}

	myPC.ping()

	myPC.duplicateRAM()
}
