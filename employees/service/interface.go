package service

import "algogrit.com/empserver/entities"

//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOPACKAGE.go

type EmployeeService interface {
	Index() ([]entities.Employee, error)
	Create(entities.Employee) (*entities.Employee, error)
}
