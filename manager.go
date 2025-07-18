package main

import "fmt"

// Manager represents an employee with leadership responsibilities
type Manager struct {
	Employee
	TeamSize int
}

// ShowInfo overrides the base method to include TeamSize
func (m Manager) ShowInfo() {
	m.Employee.ShowInfo()
	fmt.Println("Team Size:", m.TeamSize)
}
