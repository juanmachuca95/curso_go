package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // Contador

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}

	wg.Wait() // esperará la ejecución de las go rutinas
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Recibe el %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Termina")
}
