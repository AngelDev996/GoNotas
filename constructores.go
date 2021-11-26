package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

//Constructor en Go
func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	//Forma 1
	e := Employee{}
	fmt.Printf("%v\n", e)

	//Forma 2
	e2 := Employee{
		id:       1,
		name:     "Angel",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)

	//Forma 3
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 3
	e3.name = "Monika"
	e3.vacation = true

	//Forma 4
	e4 := NewEmployee(12, "Any", false)
	fmt.Printf("%v\n", *e4)
}
