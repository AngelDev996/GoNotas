//Funciones variadicas y retornos con nombre
package main

import "fmt"

func sum(values ...int) int {//esto nos permite a usar slices
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func getValues(x int) (doble int, triple int, cuadruple int) {

	doble = 2 * x
	triple = 3 * x
	cuadruple = 4 * x
	return //retornos con nombre

}

func main() {
	fmt.Println(sum(1, 5, 76))
	fmt.Println(getValues(10))
}
