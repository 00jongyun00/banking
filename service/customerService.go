package service

import (
	"bangking/domain"
	"bangking/dto"
	"bangking/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

//go:generate mockgen -destination=../mocks/service/mockCustomerService.go -package=service bangking/service CustomerService
func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	// return s.repo.FindAll(status)
	cs, err := s.repo.FindAll(status)
	response := make([]dto.CustomerResponse, 0)
	if err != nil {
		return nil, err
	}

	for _, c := range cs {
		response = append(response, c.ToDto())
	}
	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
