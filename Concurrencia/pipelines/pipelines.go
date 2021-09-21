package main

import "fmt"

func Generator(c chan<- int) { // chan <- significa que este canal solo va a escribir
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func Double(in <-chan int, out chan<- int) { // <-chan significa que este canal solo es de lectura
	for value := range in {
		out <- 2 * value // escribimos en el canal out
	}
	close(out)
}

func Print(c <-chan int) { // lectura
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	generator := make(chan int)
	double := make(chan int)

	go Generator(generator)
	go Double(generator, double)
	Print(double) // está función hay que ejecutarla sin go -- para que el programa se bloquee en esta linea.

}
