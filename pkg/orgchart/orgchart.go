package orgchart

import "github.com/KhizarShabir1/bureaucr.at/pkg/employee"

type OrgChart struct {
	CEO *employee.Employee
}

func (o *OrgChart) FindClosestCommonManager(e1 *employee.Employee, e2 *employee.Employee) *employee.Employee {
	if e1 == nil || e2 == nil {
		return nil
	}

	e1Managers := make(map[*employee.Employee]bool)
	for manager := e1.Manager; manager != nil; manager = manager.Manager {
		e1Managers[manager] = true
	}

	// Check if employee B has any of the same managers
	for manager := e2.Manager; manager != nil; manager = manager.Manager {
		if e1Managers[manager] {
			return manager
		}
	}

	return nil

}
