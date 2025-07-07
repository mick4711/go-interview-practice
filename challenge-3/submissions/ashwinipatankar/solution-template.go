package main

import (
    "fmt"
    )

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}

// AddEmployee adds a new employee to the manager's list.
func (m *Manager) AddEmployee(e Employee) {
	m.Employees = append(m.Employees, e)
}

// RemoveEmployee removes an employee by ID from the manager's list.
func (m *Manager) RemoveEmployee(id int) {
    if id < 1 {
        return 
    }
    var indexToDelete int
    var found bool
	for i := range m.Employees {
	    if m.Employees[i].ID == id {
	        indexToDelete = i
	        found = true
	        break
	    }
	}
	
	if !found {
	    return 
	}
	
	m.Employees = append(m.Employees[:indexToDelete], m.Employees[indexToDelete+1:]...)
}

// GetAverageSalary calculates the average salary of all employees.
func (m *Manager) GetAverageSalary() float64 {
	var avgSalary float64
	
	if len(m.Employees) < 1 {
	    return avgSalary
	}
	
	for i := range m.Employees {
	    avgSalary += m.Employees[i].Salary
	}
	
	
	return avgSalary/float64(len(m.Employees))
}

// FindEmployeeByID finds and returns an employee by their ID.
func (m *Manager) FindEmployeeByID(id int) *Employee {
	for i := range m.Employees {
	    if m.Employees[i].ID == id {
	        return &m.Employees[i]
	    }
	}
	
	return nil
}

func main() {
	manager := Manager{}
	manager.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	manager.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
	manager.RemoveEmployee(1)
	averageSalary := manager.GetAverageSalary()
	employee := manager.FindEmployeeByID(2)

	fmt.Printf("Average Salary: %f\n", averageSalary)
	if employee != nil {
		fmt.Printf("Employee found: %+v\n", *employee)
	}
}
