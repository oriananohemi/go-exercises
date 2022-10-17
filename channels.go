package channels

import "fmt"

// Channels es una forma de organizar las goroutines. Te permite manejar datos entre goroutines. Es un conducto donde se puede transportar solo 1 tipo de dato

// Deberiamos indicarle si va a recibir es de entrada o salida de datos (c <-)
func say(text string, c <-chan string) {
	/* 	c <- text */
	text = <-c
}

func main() {
	// Se debe indicar el tipo de datos
	// Opcional se le indica cuantos datos simultaneos va a tomar, sino se envia va a tomar un dato dinamico
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)
}
