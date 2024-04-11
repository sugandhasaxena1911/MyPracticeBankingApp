package repository

import "github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/core/domain"

type CustomerRepositoryStub struct {
}

func (crs *CustomerRepositoryStub) GetAllCustomers() ([]domain.Customers, error) {

	custs := []domain.Customers{
		{Custname: "Sugandha", Custid: "1", Custcity: "Lucknow", Custbirthdate: "2023-11-05", Custzipcode: "6766", Custstatus: "Active"},
	}

	return custs, nil
}

func NewCustomerReposioryStub() *CustomerRepositoryStub {
	return &CustomerRepositoryStub{}
}
