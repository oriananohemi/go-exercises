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

	// tipos
	// Si no se especifica el tipo, toma por defecto el del computador, ejemplo "int16 = 16bits = -2^15 a 2^15-1"
	int     // Depende del OS (32 o 64 bits)
	uint    // Números positivos
	float   // Números decimales (se puede cambiar entre 32 o 64)
	string  // DEBE IR CON "" comillas dobles
	complex // hay de 32 y 64, se guardan numeros complejos

	// fmt para imprimir
	// https: //pkg.go.dev/fmt
	fmt.Println("Hola mundo")

	// como buena practica si tenemos argumentos del mismo tipo unimos los argumentos. Ejemplo a,b int
	func tripleArgument(a,b int, c string) {
		fmt.Println(a,b,c)
	}

	func doubleReturn(a int) (c,d int) {
		return a, a *2
	}

	value1, value2 := doubleReturn(2)

	// ciclos
	for i := 0; i<= 10; i++ {
		fmt.Println((i))
	}

	// For while
	counter := 0
	for i := 0; i<= 10; i++ {
		fmt.Println(counter)
		counter++
	}

	// For forever - No tiene condicion
	counter := 0
	for {
		fmt.Println(counter)
		counter++
	}

}

// go build "main" --- compila el archivo del código
// go run "main" --- crea archivo temporal
