//En go no existe la herencia como tal, el polimorfirsmo no funciona como en java, para esto ultimo es necesario la implementacion de interfaces
//Ya que no hay herencia como tal, aplicamos el conceptp de composicion
package main

import "fmt"

type Persona struct {
	name string
	edad int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Persona
	Employee
}

func GetMessage(p Persona) {
	fmt.Printf("%v with age %d\n", p.name, p.edad)
}

func main() {
	fte := FullTimeEmployee{}
	fte.edad = 23
	fte.id = 1
	fte.name = "Monika"
	fmt.Printf("%v\n", fte)
}
