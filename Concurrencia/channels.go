package main

import "fmt"

/**
* Tipos de canales
* Canales con buffer y si buffer
 */

func function() {
	/* Primer bloque
	c := make(chan int)

	c <- 1

	fmt.Println(<-c)

	al ejecutar go run channels.go esto devolverá un error
	fatal error: all goroutines are asleep - deadlock!

	goroutine 1 [chan send]:
	main.main()
			/home/juan/Escritorio/go/curso_go/Concurrencia/channels.go:13 +0x59
	exit status 2

	ya que necesitamos recibir lo que envia el canal

	*/

	c := make(chan int, 3) // channel with buffer

	c <- 1
	c <- 5
	c <- 5
	// c <- 2 Mi buffer es de 3 si envio más de tres retornará error

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
