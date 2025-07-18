package main

func main() {
	// Developer and Manager instances
	dev := Developer{
		Employee:  NewEmployee("Amanda Limón", 26, "Frontend Developer", 25000),
		TechStack: "React, TypeScript",
	}

	mgr := Manager{
		Employee: NewEmployee("Carolina Fernández", 28, "Project Manager", 38000),
		TeamSize: 6,
	}

	// Display initial information
	dev.ShowInfo()
	mgr.ShowInfo()

	// Increase their salaries
	dev.IncreaseSalary(10)
	mgr.IncreaseSalary(5)

	// Display updated information
	dev.ShowInfo()
	mgr.ShowInfo()
}
