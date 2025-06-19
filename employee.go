package main

import (
	"fmt"
	"strings"
)

// Employee represents a basic employee with attributes
type Employee struct {
	Name     string
	Age      int
	Position string
	Salary   float64
}

// NewEmployee acts as a constructor to create a new Employee
func NewEmployee(name string, age int, position string, salary float64) Employee {
	return Employee{
		Name:     name,
		Age:      age,
		Position: position,
		Salary:   salary,
	}
}

// formatWithCommas formats a float64 value with commas and two decimal places
func formatWithCommas(value float64) string {
	parts := strings.Split(fmt.Sprintf("%.2f", value), ".")
	intPart := parts[0]
	decimalPart := parts[1]

	var result strings.Builder
	n := len(intPart)
	for i, digit := range intPart {
		if i > 0 && (n-i)%3 == 0 {
			result.WriteRune(',')
		}
		result.WriteRune(digit)
	}

	return result.String() + "." + decimalPart
}

// ShowInfo prints the employee's details
func (e Employee) ShowInfo() {
	fmt.Println("------ Employee Information ------")
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
	fmt.Println("Position:", e.Position)
	fmt.Printf("Salary: $%s\n", formatWithCommas(e.Salary))
}

// IncreaseSalary increases the salary by a given percentage
func (e *Employee) IncreaseSalary(percentage float64) {
	increase := e.Salary * percentage / 100
	e.Salary += increase
	fmt.Printf("Salary increased by %.2f%% ($%s)\n", percentage, formatWithCommas(increase))
}
