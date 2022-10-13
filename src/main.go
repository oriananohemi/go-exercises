package main

import "fmt"

func main() {
	// Declaración
	const pi float64 = 3.1416
	const pi2 = 3.1416
	base := 5 // declara y asigna la variable
	var altura int = 14
	var rango int

	// Las variables tienen un valor por defecto, no se asigna null
	var a int     // 0
	var b float64 // 0
	var c string  // " "
	var d bool    // false

	fmt.Println(a, b, c, d)

	fmt.Println("Hola mundo")

}

// go build "main" --- compila el archivo del código
// go run "main" --- crea archivo temporal
