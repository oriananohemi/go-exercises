package interfaceType

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func calcular(f figuras2D) {
	fmt.Println("Area", f.area())
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func getArea() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	calcular(myCuadrado)
	calcular(myRectangulo)
}
