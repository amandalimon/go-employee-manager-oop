package main

import "fmt"

// Developer represents an employee with technical expertise
type Developer struct {
	Employee
	TechStack string
}

// ShowInfo overrides the base method to include TechStack
func (d Developer) ShowInfo() {
	d.Employee.ShowInfo()
	fmt.Println("Tech Stack:", d.TechStack)
}
