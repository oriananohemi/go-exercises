package puntero

import "fmt"

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
