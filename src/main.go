package main

import (
	"fmt"
	"strconv"
)

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

	// Negacion !
	myVar := true
	!myVar


	// Convertir texto a numero
	value, err := strconv.Atoi("53")

	// defer, es lo último que se ejecuta, es para cerrar una conexion a base de datos, algo que se debe cerrar.
	defer fmt.Println(("hola"))

	// ARRAY INMUTABLE - SLICE MUTABLE
	var array[4]int
	// tamaño
	len(array)
	// capacidad
	cap(array)

	// La diferencia con Array es que a este no se le indica cuantos elementos va a tener
	slice := []int{0,1,2,3,4}

	//el ultimo numero es exclusivo, es decir, no se toma en cuenta, el primero es inclusivo, ese si se toma en cuenta.
	slice[:3]
	slice[2:4]


}

// go build "main" --- compila el archivo del código
// go run "main" --- crea archivo temporal
