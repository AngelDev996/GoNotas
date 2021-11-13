package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup

	// say("Hola")
	wg.Add(1)
	go say("Mundo", &wg)
	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("adios")
	time.Sleep(time.Second * 1)

}
