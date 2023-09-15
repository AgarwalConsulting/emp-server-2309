package repository

import "algogrit.com/empserver/entities"

//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOPACKAGE.go

type EmployeeRepository interface {
	ListAll() ([]entities.Employee, error)
	Save(newEmp entities.Employee) (*entities.Employee, error)
}
