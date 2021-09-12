package main

import "fmt"

/*Struct es un equivalente a una class en programación orientada a objectos*/
type Employee struct {
	id   int
	name string
}

/* Functions Receiver and Pointer receiver*/
func (e *Employee) SetId(id int) {
	e.id = id
}

func main() {
	e := Employee{}

	fmt.Printf("%v\n", e) // Go asigna valores al crear, ej: id 0, "" -> retornara {0 }
	e.id = 1
	e.name = "Juan Gabriel"

	e.SetId(1995)
	fmt.Printf("%v\n", e)
}
