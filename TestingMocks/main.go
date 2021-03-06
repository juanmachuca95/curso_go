package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

var GetPersonByDNI = func(_ string) (Person, error) { // _ = dni
	time.Sleep(5 * time.Second)
	return Person{}, nil
}

var GetEmployeeById = func(_ int) (Employee, error) { // _ = id
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployeeById(id)
	if err != nil {
		return FullTimeEmployee{}, err
	}
	ftEmployee.Employee = e

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return FullTimeEmployee{}, nil
	}
	ftEmployee.Person = p

	return ftEmployee, nil
}
