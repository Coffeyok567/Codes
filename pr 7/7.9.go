// Корчагин Евгений 363
package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

type Department struct {
	Name      string
	employees map[int]*Employee
}

func (d *Department) AddEmployee(newEmployee *Employee) {
	if d.employees == nil {
		d.employees = make(map[int]*Employee)
	}
	fmt.Printf("%s нанят, отдел %s\n", newEmployee.Name, d.Name)
	d.employees[newEmployee.ID] = newEmployee
}

func (d *Department) RemoveEmployee(employeeID int) error {
	employee, ok := d.employees[employeeID]
	if ok {
		fmt.Printf("%s уволен, отдел %s\n", employee.Name, d.Name)
		delete(d.employees, employeeID)
		return nil
	} else {
		return errors.New("Не найден")
	}
}

func (d *Department) CalculateSalaryFund() float64 {
	var sum float64
	for employee := range d.employees {
		sum += d.employees[employee].Salary
	}
	return sum
}

func (d *Department) GetEmployeesByPosition(position string) []*Employee {
	var result []*Employee
	for employee := range d.employees {
		if d.employees[employee].Position == position {
			result = append(result, d.employees[employee])
		}
	}
	return result
}

func main() {
	hrDepartment := Department{Name: "HR"}

	emp1 := &Employee{ID: 1, Name: "JDH", Position: "Разработчик", Salary: 50000}
	emp2 := &Employee{ID: 2, Name: "Дилан", Position: "Разработчик", Salary: 80000}
	emp3 := &Employee{ID: 3, Name: "lololowka", Position: "Менеджер", Salary: 55000}

	hrDepartment.AddEmployee(emp1)
	hrDepartment.AddEmployee(emp2)
	hrDepartment.AddEmployee(emp3)

	fund := hrDepartment.CalculateSalaryFund()
	fmt.Printf("Зарплатный фонд: %.2f\n", fund)

	managers := hrDepartment.GetEmployeesByPosition("Менеджер")
	fmt.Println("Менеджеры:")
	for _, manager := range managers {
		fmt.Printf("%s\n", manager.Name)
	}

	err := hrDepartment.RemoveEmployee(2)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	newFund := hrDepartment.CalculateSalaryFund()
	fmt.Printf("новый зарплатный фонд: %.2f\n", newFund)
}
