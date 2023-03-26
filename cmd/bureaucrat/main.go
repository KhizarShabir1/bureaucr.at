package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/KhizarShabir1/bureaucr.at/pkg/employee"
	"github.com/KhizarShabir1/bureaucr.at/pkg/orgchart"
)

func main() {

	// Create a sample organization chart
	ceo := &employee.Employee{ID: 1, Name: "Claire"}
	manager1 := &employee.Employee{ID: 2, Name: "Manager 1", Manager: ceo}
	manager2 := &employee.Employee{ID: 3, Name: "Manager 2", Manager: ceo}

	employee1 := &employee.Employee{ID: 4, Name: "Employeee 1", Manager: manager1}
	employee2 := &employee.Employee{ID: 5, Name: "Employeee 2", Manager: manager1}
	employee3 := &employee.Employee{ID: 6, Name: "Employeee 3", Manager: manager2}
	employee4 := &employee.Employee{ID: 7, Name: "Employeee 4", Manager: manager2}

	ceo.DirectReports = []*employee.Employee{manager1, manager2}

	manager1.DirectReports = []*employee.Employee{employee1, employee2}
	manager2.DirectReports = []*employee.Employee{employee3, employee4}

	org := &orgchart.OrgChart{CEO: ceo}

	//pasring command line argguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: cmd/bureaucrat/main.go employee-id-1 employee-id-2")
		os.Exit(1)
	}

	id1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid employee ID: ", os.Args[1])
		os.Exit(1)
	}

	id2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid employee ID: ", os.Args[2])
		os.Exit(1)
	}
	// use these two ids to find the employees
	emp1 := findEmployeeById(ceo, id1)
	emp2 := findEmployeeById(ceo, id2)

	manager := org.FindClosestCommonManager(emp1, emp2)

	if manager != nil {
		fmt.Println("Closest common manager is,", manager.Name)
	} else {
		fmt.Println("No common manager found")
	}
}

func findEmployeeById(root *employee.Employee, id int) *employee.Employee {
	if root == nil {
		return nil
	}
	if root.ID == id {
		return root
	}
	for _, report := range root.DirectReports {
		result := findEmployeeById(report, id)
		if result != nil {
			return result
		}
	}
	return nil
}
