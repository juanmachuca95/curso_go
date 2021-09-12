package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{ // Importante el & para retornar la referencia al objeto que estamos creando
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {

	// Forma 1 de usar constructor por defecto golang
	e := Employee{}
	fmt.Printf("Forma 1°: %v\n", e)

	// Forma 2 de usar constructor asignando valores
	e2 := Employee{
		id:       1,
		name:     "Juan Gabriel",
		vacation: false,
	}
	fmt.Printf("Forma 2°: %v\n", e2)

	// Forma 3 de usar constructor
	e3 := new(Employee)              // Importante tener en cuenta que esto devuelve la referencia al objecto
	fmt.Printf("Forma 3°: %v\n", e3) // retorna &{0  false}
	fmt.Printf("Forma 3° Correcta: %v\n", *e3)

	// Forma 4 - Forma ganadora
	e4 := NewEmployee(1995, "Juan Gabriel Machuca", true)
	fmt.Printf("Forma 4°: %v\n", e4)             // Se imprime la referencia al objeto
	fmt.Printf("Forma 4° (Correcta): %v\n", *e4) // Se imprime los valores del objeto referenciado

}
