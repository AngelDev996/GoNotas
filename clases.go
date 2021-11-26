package main

import "fmt"

type Employee struct { // esto es lo equivalente a una clase en java
	id     int
	nombre string
}

//Funcion main
func main() {
	e := Employee{} //Asi instanciamos los objetos
	e.id = 4
	e.nombre = "Angel"
	fmt.Printf("%v", e)

}
