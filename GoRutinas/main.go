package main

import (
	"fmt"
	"time"
)

func main() {

	mapita := map[string]int{ // Lo que hara mi main mientras tanto
		"uno": 1, "dos": 2, "tres": 3,
	}

	for k, v := range mapita {
		fmt.Printf("Mapita[index] : %s - Mapita[value]: %d\n", k, v)
	}

	channel := make(chan int) // canal que envia y devuel un entero
	go miRutina(channel)
	x := <-channel // Recibe
	fmt.Println(x)
}

func miRutina(channel chan int) {
	fmt.Println("Voy a dormir  3 segundos . . .")
	time.Sleep(3 * time.Second)
	fmt.Println("He terminado de dormir!!")
	channel <- 1
}
