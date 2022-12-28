package service

import (
	"github.com/Gustavo-Zamai/Banking-API/domain"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
    repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(Id string) (*domain.Customer, error) {
	return s.repo.ById(Id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}

