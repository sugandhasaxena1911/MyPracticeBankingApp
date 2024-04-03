package service

import (
	"fmt"
	"log"

	"github.com/sugandhasaxena19/MyPracticeBankingApp/helpers/error"
	"github.com/sugandhasaxena19/MyPracticeBankingApp/internal/core/domain"
	"github.com/sugandhasaxena19/MyPracticeBankingApp/internal/core/dto"
	"github.com/sugandhasaxena19/MyPracticeBankingApp/internal/core/ports/repositories"
)

const (
	ACTIVE   = "active"
	INACTIVE = "inactive"
)

type CustomerServiceDefault struct {
	CustRepo repositories.CustomerRepository
}

func (csd *CustomerServiceDefault) FindAllCustomers() ([]dto.CustomerDTO, *error.AppError) {
	log.Println("Inside Find all customers ")
	customers, err := csd.CustRepo.GetAllCustomers()
	var custdto []dto.CustomerDTO
	if err != nil {
		return custdto, err
	}
	custs := []dto.CustomerDTO{}
	for _, v := range customers {
		c := v.ToCustomerDTO()
		custs = append(custs, c)

	}
	return custs, nil
}

func validateStatus(status string) *error.AppError {
	if status == ACTIVE || status == INACTIVE {
		return nil
	}

	return error.NewBadRequestAppError(fmt.Sprintf("invalid status :%s", status))
}

func (csd *CustomerServiceDefault) FindAllCustomersByStatus(status string) ([]dto.CustomerDTO, *error.AppError) {
	err := validateStatus(status)
	if err != nil {
		return nil, err
	}
	custs, err := csd.CustRepo.GetAllCustomersByStatus(status)
	if err != nil {
		return nil, err
	}
	dtocusts := []dto.CustomerDTO{}
	for _, c := range custs {
		dtocust := c.ToCustomerDTO()
		dtocusts = append(dtocusts, dtocust)
	}

	return dtocusts, nil
}

func (csd *CustomerServiceDefault) FindCustomerByID(custid string) (dto.CustomerDTO, *error.AppError) {
	log.Println("Inside customerservice ", custid)
	customer, err := csd.CustRepo.GetCustomerByID(custid)
	log.Println("after customer repo db  ", err)

	var custdto dto.CustomerDTO
	if err != nil {
		return custdto, err
	}
	custdto = customer.ToCustomerDTO()
	return custdto, err
}

func (csd *CustomerServiceDefault) CreateCustomer(cust dto.CustomerDTO) (dto.CustomerDTO, *error.AppError) {
	domaincust := domain.Customers{}
	status := cust.Custstatus

	err := validateStatus(status)
	if err != nil {
		return dto.CustomerDTO(domaincust), err
	}
	switch status {
	case "active":
		status = "1"
	case "inactive":
		status = "0"
	}

	domaincust = domain.Customers{
		Custname:      cust.Custname,
		Custbirthdate: cust.Custbirthdate,
		Custcity:      cust.Custcity,
		Custzipcode:   cust.Custzipcode,
		Custstatus:    status}
	domaincust, err = csd.CustRepo.PostCustomer(domaincust)
	if err != nil {
		return cust, err
	}

	return domaincust.ToCustomerDTO(), nil

}

func NewCustomerServiceDefault(custrepository repositories.CustomerRepository) *CustomerServiceDefault {
	return &CustomerServiceDefault{custrepository}

}

/*
func NewCustomerServiceDB() *CustomerServiceDefault {
	//return &CustomerServiceDefault{repository.NewCustomerReposioryStub()}
	return &CustomerServiceDefault{repository.NewCustomerReposioryDB()}

}*/
