package service

import (
	_ "github.com/golang/mock/mockgen/model"
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/core/dto"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/service/mockCustomerService.go -package=service github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/core/ports/service CustomerService
type CustomerService interface {
	FindAllCustomers() ([]dto.CustomerDTO, *error.AppError)
	FindAllCustomersByStatus(status string) ([]dto.CustomerDTO, *error.AppError)
	FindCustomerByID(custid string) (dto.CustomerDTO, *error.AppError)
	CreateCustomer(dto.CustomerDTO) (dto.CustomerDTO, *error.AppError)
}
