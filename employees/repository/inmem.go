package repository

import "algogrit.com/empserver/entities"

type inMemRepo struct {
	employees []entities.Employee
}

func (repo *inMemRepo) ListAll() ([]entities.Employee, error) {
	return repo.employees, nil
}

func (repo *inMemRepo) Save(newEmp entities.Employee) (*entities.Employee, error) {
	newEmp.ID = len(repo.employees) + 1

	repo.employees = append(repo.employees, newEmp)

	return &newEmp, nil
}

func NewInMem() EmployeeRepository {
	var employees = []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
		{2, "Anupam", "Cloud", 10002},
		{3, "Udbhav", "SRE", 20002},
	}

	return &inMemRepo{employees}
}
