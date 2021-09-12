package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func sumArrayParam(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

// Funcion con nombres - Go infiere la suma de los int
func getValues(x int) (p int, s int, t int) {
	p = x * 2
	s = x * 3
	t = x * 4
	return
}

func main() {
	x := 20
	/* Funcion anonima */
	y := func() int {
		return x * 123
	}() // Si no colocamos () la función no se ejecuta

	fmt.Printf("El valor de y es : %d\n", y)

	fmt.Printf("Suma de dos números: %d\n", sum(1, 2))

	fmt.Printf("Suma con número de parametros indefinido: %d\n", sumArrayParam(1, 2, 3, 4, 5, 10))

	fmt.Println(getValues(2))

}
