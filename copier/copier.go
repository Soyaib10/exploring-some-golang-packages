package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type Employee struct {
	Name   string
	Salary *int
}

type Payroll struct {
	Name   string
	Salary *int
}

func CopierFunction() {
	initialSalary := 5000
	employee := Employee{
		Name:   "Soyaib",
		Salary: &initialSalary,
	}

	var payroll Payroll

	// copier performs a deep copy for pointers by default
	copier.Copy(&payroll, &employee)
	fmt.Printf("Employee Salary: %d\n", *employee.Salary)
	fmt.Printf("Payroll Salary: %d\n", *payroll.Salary)

	// Changing the value in the source
	*employee.Salary = 7000

	fmt.Printf("Employee Salary: %d\n", *employee.Salary)
	fmt.Printf("Payroll Salary: %d\n", *payroll.Salary)

	// Checking if they point to the same memory address
	fmt.Printf("Employee Salary Address: %p\n", employee.Salary)
	fmt.Printf("Payroll Salary Address: %p\n", payroll.Salary)
}

func CopierTags() {
	type A struct {
		Name   string
		secret int // unexported values 
		Age    int 
	}

	type B struct {
		Name   string
		secret int
		Age    int `copier:"-"`
	}

	a := A{
		Name:   "John Doe",
		secret: 3456,
		Age:    30,
	}

	b := B{}
	copier.Copy(&b, &a)
	fmt.Printf("B: %#v\n", b)
}
