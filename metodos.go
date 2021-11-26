package main

import "fmt"

type Employee struct { // esto es lo equivalente a una clase en java
	id     int
	nombre string
}

//Asi declaramos metodos en go (Setter en este caso)
func (e *Employee) SetId(id int) {
	e.id = id
}

//Asi declaramos metodos en go (Setter en este caso)
func (e *Employee) SetName(name string) {
	e.nombre = name
}

//getter
func (e *Employee) GetName() string {
	return e.nombre
}

//getter
func (e *Employee) GetId() int {
	return e.id
}

//Funcion main
func main() {
	e := Employee{} //Asi instanciamos los objetos
	e.id = 4
	e.nombre = "Angel"
	fmt.Printf("%v", e)
	e.SetId(12)
	e.SetName("Monika")
	fmt.Printf("%v", e)
}
