package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person // Importante poner Person -> como propiedad anonima
	Employee
	salary float64
}

func GetMessage(p *Person) {
	fmt.Printf("El nombre es %s y su edad es %d\n", p.name, p.age)
}

func main() {

	fullTimeEmployee := FullTimeEmployee{}
	fmt.Printf("FullTimeEmployee: %v\n", fullTimeEmployee) // Se podrá ver el objeto segmentado

	/* Herencia aplicada*/
	fullTimeEmployee.id = 1
	fullTimeEmployee.age = 26
	fullTimeEmployee.name = "Juan Gabriel Machuca"
	fullTimeEmployee.salary = 26000.00

	fmt.Printf("FullTimeEmployee: %v\n", fullTimeEmployee) // Se podrá ver el objeto segmentad

	/* Si quisieramos aplicar polimorfismo  -> no funcionaria. Para ello ver Interfaces */
	//GetMessage(fullTimeEmployee) // RETORNARA ERROR: no puede usar FullTimeEmployee como una Person.
}
