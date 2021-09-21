package main

import (
	"fmt"
	"time"
)

/**
* Multiplexación - palabra reservada select
 */

func main() {
	channelOne := make(chan int)
	channelTwo := make(chan int)

	go doSomething(10, channelOne, 10)
	go doSomething(2, channelTwo, 2)

	/* fmt.Println(<-channelOne)
	fmt.Println(<-channelTwo) */

	for i := 0; i < 2; i++ {
		select {
		case msgChannelOne := <-channelOne:
			fmt.Printf("Se ha terminado de ejecutar el channel ONE duración de %d segundos\n", msgChannelOne)

		case msgChannelTwo := <-channelTwo:
			fmt.Printf("Se ha terminado de ejecutar el channel TWO duración de %d segundos\n", msgChannelTwo)
		}
	}
}

func doSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i * time.Second)
	c <- param
}
