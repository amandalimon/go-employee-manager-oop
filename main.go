package main

func main() {
	// Create two employees using the constructor
	employee1 := NewEmployee("Amanda Limón", 26, "Frontend Developer", 25000)
	employee2 := NewEmployee("Carolina Fernández", 28, "Project Manager", 38000)

	// Display initial information
	employee1.ShowInfo()
	employee2.ShowInfo()

	// Increase their salaries
	employee1.IncreaseSalary(10)
	employee2.IncreaseSalary(5)

	// Display updated information
	employee1.ShowInfo()
	employee2.ShowInfo()
}
