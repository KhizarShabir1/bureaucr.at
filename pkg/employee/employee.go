package employee

type Employee struct {
	ID            int
	Name          string
	Manager       *Employee
	DirectReports []*Employee
}

func (e *Employee) IsManager() bool {
	return len(e.DirectReports) > 0
}
