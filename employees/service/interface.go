package service

import "algogrit.com/empserver/entities"

type EmployeeService interface {
	Index() ([]entities.Employee, error)
	Create(entities.Employee) (*entities.Employee, error)
}
