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
	Person
	Employee
	salary float64
}

type TemporaryEmployee struct {
	Person
	Employee
	dateEnd string
}

type TypeEmployee interface { // Interface
	getInfoEmployee() string // Contrato
}

func getInfoEmployee(typeEmployee TypeEmployee) { // Como se resuelve la interface
	fmt.Println(typeEmployee.getInfoEmployee())
}

func (tEmployeeFulltime FullTimeEmployee) getInfoEmployee() string {
	name := tEmployeeFulltime.name
	age := tEmployeeFulltime.age
	salary := tEmployeeFulltime.salary
	info := fmt.Sprintf("El empleado: %s - edad: %d - salario: %2.f es empleado de tiempo completo", name, age, salary)
	return info
}

func (tEmployeeTemporary TemporaryEmployee) getInfoEmployee() string {
	name := tEmployeeTemporary.name
	age := tEmployeeTemporary.age
	dateEnd := tEmployeeTemporary.dateEnd
	info := fmt.Sprintf("El empleado: %s - edad: %d - finaliza: %s es empleado temporario", name, age, dateEnd)
	return info
}

func main() {

	tEmployeeFullTime := FullTimeEmployee{}
	tEmployeeFullTime.name = "Maria Laura Torres"
	tEmployeeFullTime.age = 25
	tEmployeeFullTime.salary = 10000.00

	tEmployeeTemporary := TemporaryEmployee{}
	tEmployeeTemporary.name = "Juan Gabriel Machuca"
	tEmployeeTemporary.age = 25
	tEmployeeTemporary.dateEnd = "18 de Enero 2022"

	/*Poliformismo*/
	getInfoEmployee(tEmployeeFullTime)
	getInfoEmployee(tEmployeeTemporary)

}
