package main

import "fmt"

/*Struct es un equivalente a una class en programación orientada a objectos*/
type Employee struct {
	id   int
	name string
}

/* Functions Receiver and Pointer receivers*/
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}

	fmt.Printf("%v\n", e) // Go asigna valores al crear, ej: id 0, "" -> retornara {0 }

	e.id = 1
	e.name = "Juan Gabriel"
	fmt.Printf("%v\n", e)

	fmt.Println("Funciones Set de Employee")
	e.SetId(1995) // Importante la forma en la que golang interpreta la función definida arriba. (Pointer receiver)
	fmt.Printf("%v\n", e)

	e.SetName("Juan Golang")
	fmt.Printf("%v\n", e)

	// Mostrar datos del function get
	fmt.Println("Funciones Get de Employee")
	fmt.Printf("GetId: %d\n", e.GetId())
	fmt.Printf("GetName: %s\n", e.GetName())
}
