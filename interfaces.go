package main

import "fmt"

//Interfaces
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

func (r rectangulo) area() float64 {
	return r.altura * r.base
}

func (c cuadrado) area() float64 {

	return c.base * c.base
}

func calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}
func main() {

	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}
	fmt.Println(myCuadrado)
	fmt.Println(myRectangulo)
	calcular(myCuadrado)
	calcular(myRectangulo)

}
