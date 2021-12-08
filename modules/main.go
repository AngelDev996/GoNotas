//go mod init github.com/AngelDev996/test -> sintaxis general para crear modulos personalizados
//go get github.com/donvito/hellomod -> sintaxis para clonar un modulo de github
//go mod tidy elimina dependencias no utilizadas
//go buil main.go compila el codigo
//go mod download -json
package main

import (
	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2" // usar alias porque ambos paquetes se llaman igual
)

func main() {
	hellomod.SayHello()
	hellomod2.SayHello("Monika") //alias del otro modulo
}
