package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 8
	y := 7
	fmt.Printf("Variable x: %d\n", x)
	fmt.Printf("Variable y: %d\n", y)

	dato, err := strconv.ParseInt("10", 10, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(dato)

	/* Mapas */
	mapa := make(map[string]int)
	mapa["Llave"] = 16
	fmt.Printf("Mapa (key: Llave) (valor: %d)\n", mapa["Llave"])

	/* Slices */
	slices := []int{1, 2, 3}
	for index, value := range slices {
		fmt.Printf("Indice: %d - Valor: %d\n", index, value)
	}

	slices = append(slices, 100) // Agregar elemento a un slice
	for index, value := range slices {
		fmt.Printf("Indice: %d - Valor: %d\n", index, value)
	}

}
