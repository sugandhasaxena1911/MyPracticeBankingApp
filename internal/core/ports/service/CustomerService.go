package service

import (
	_ "github.com/golang/mock/mockgen/model"
<<<<<<< HEAD
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/core/dto"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/service/mockCustomerService.go -package=service github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/core/ports/service CustomerService
=======
	"github.com/sugandhasaxena19/MyPracticeBankingApp/helpers/error"
	"github.com/sugandhasaxena19/MyPracticeBankingApp/internal/core/dto"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/service/mockCustomerService.go -package=service github.com/sugandhasaxena19/MyPracticeBankingApp/internal/core/ports/service CustomerService
>>>>>>> 27c2bab4ae9973b95478eedee3812fd6c4b17ae0
type CustomerService interface {
	FindAllCustomers() ([]dto.CustomerDTO, *error.AppError)
	FindAllCustomersByStatus(status string) ([]dto.CustomerDTO, *error.AppError)
	FindCustomerByID(custid string) (dto.CustomerDTO, *error.AppError)
	CreateCustomer(dto.CustomerDTO) (dto.CustomerDTO, *error.AppError)
}
