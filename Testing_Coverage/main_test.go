package main

import "testing"

/* El comando de go para ejecutar test es:
*  go test
 */

/*
	Para obtener una idea de cuan testeado esta nuestro codigo
	podemos ejecutar:
	go test -cover
*/

/*
	Importante ejecutar este comando cada vez que actualizemos los test para que nuestro archivo
	tenga los ultimos en las Metricas: 

	go test -coverprofile=coverago.out
	
	creará un archivo para obtener metricas sobre los test en nuestra aplicación ejecutar:

	go tool cover -func=coverage.out
	Ejemplo:
	github.com/juanmachuca95/curso_go/Testing_Coverage/main.go:5:   Sum             100.0% // Super testeado
	github.com/juanmachuca95/curso_go/Testing_Coverage/main.go:9:   GetMax          0.0% // Nada testeado
	github.com/juanmachuca95/curso_go/Testing_Coverage/main.go:16:  main            0.0%
	total:                                                          (statements)    50.0%


	Este comando nos mostrará un html con información aún más especifica
	go tool cover -html=coverage.out

*/
func TestSum(t *testing.T) {
	total := Sum(1, 2)

	if total != 3 {
		t.Errorf("El valor es incorrecto, se obtuvo %d y se esperaba %d\n", total, 3)
	}
}

 
func TestSum2(t *testing.T) {
	tables := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3}, // El 3 (c) es el resultado de a + b (1 + 2)
		{5, 5, 10},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)

		if total != item.c {
			t.Errorf("El valor es incorrecto, se obtuvo %d y se esperaba %d\n", total, item.c)
		}
	}
}
 
func TestGetMax(t *testing.T) { // DEbes escribir bien Test --> si pones Text no funciona
	tables := []struct {
		x int
		y int
		r int
	}{
		{4, 2, 4},
		{7, 1, 7},
		{9, 8, 9},
		{3, 10, 10}, // si se cubren todos los casos habrá un coverage del 100%
	}

	for _, item := range tables {
		max := GetMax(item.x, item.y)
		if max != item.r {
			t.Errorf("El test ha fallado, se esperaba %d pero se obtuvo %d\n", item.r, max)
		}
	}
}
