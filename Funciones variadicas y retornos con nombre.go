//Funciones variadicas y retornos con nombre
/*
Funciones variadicas y retornos con nombre.

Las funciones variadicas nos permiten utilizar como slices los argumentos de funciones de los cuales no sabemos su longitud exacta.

Los retornos con nombre nos permiten definir variables antes de definir el cuerpo de la función, por lo cual utilizaremos return para devolverlos.
*/
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
