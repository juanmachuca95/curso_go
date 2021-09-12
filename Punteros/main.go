package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)

	/* &obtiene la referencia a la variable */
	y := &x        // y va ser un apuntador a x
	fmt.Println(y) // Devolver una dirección de memoria 0xc0000140c0 por ej.

	/* *obtiene el valor en la referencia */

	fmt.Println(*y) //Si quisieramos ver el valor de y entonces se haria esto
}
