//Uso de interfaces para simular polimorfismo
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
	endDate string
}

type TemporaryEmployee struct {
	Persona
	Employee
	taxRate int
}
type PrintInfo interface {
	getMessage() string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	fte := FullTimeEmployee{}
	te := TemporaryEmployee{}
	getMessage(fte)
	getMessage(te)
}
