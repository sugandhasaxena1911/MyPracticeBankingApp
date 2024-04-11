package repositories

import (
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/core/domain"
)

type CustomerRepository interface {
	GetAllCustomers() ([]domain.Customers, *error.AppError)
	GetAllCustomersByStatus(status string) ([]domain.Customers, *error.AppError)
	GetCustomerByID(custid string) (domain.Customers, *error.AppError)
	PostCustomer(domain.Customers) (domain.Customers, *error.AppError)
}
