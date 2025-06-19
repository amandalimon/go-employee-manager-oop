# OOP Employee Management in Go

This repository contains a simple academic project to demonstrate object-oriented programming concepts using the Go programming language. The project simulates a basic employee management system with data structures, methods, and constructors.

## 🧱 Project Structure

```
.
├── main.go          // Main application entry point
└── employee.go      // Struct definition and methods
```

## 🚀 How to Run the Project

### ✅ Prerequisites

Go installed (version 1.18+ recommended)  
📥 [Download from go.dev](https://go.dev/dl)  

---
### ▶️ Running the Application

#### 1. **Clone the repository**

```bash
git clone https://github.com/amandalimon/go-employee-manager-oop.git
```

Then navigate to the project directory:

```bash
cd go-employee-manager-oop
```

#### 2. **Run the application using the Go CLI**

```bash
go run main.go employee.go
```

You should see the employee data printed and salary adjustments applied via methods.

---

### 💡 What should you expect?

When the program runs, it will:

1. Create two `Employee` instances using the constructor function `NewEmployee`.
2. Display the initial information of both employees using the `ShowInfo` method.
3. Apply a salary increase to each employee using the `IncreaseSalary` method.
4. Display the updated information showing the new salaries.

The output will look similar to:

```
------ Employee Information ------
Name: Amanda Limón
Age: 26
Position: Frontend Developer
Salary: $25,000.00

------ Employee Information ------
Name: Carolina Fernández
Age: 28
Position: Project Manager
Salary: $38,000.00
Salary increased by 10.00% ($2,500.00)
Salary increased by 5.00% ($1,900.00)

------ Employee Information ------
Name: Amanda Limón
Age: 26
Position: Frontend Developer
Salary: $27,500.00

------ Employee Information ------
Name: Carolina Fernández
Age: 28
Position: Project Manager
Salary: $39,900.00
```

### ❗ Common issues

- If you see `command not found: go`, make sure Go is correctly installed and added to your system’s PATH.
- Always use the correct file names: `main.go` and `employee.go` (case-sensitive on some systems).
