package repository

<<<<<<< HEAD
import "github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/core/domain"
=======
import "github.com/sugandhasaxena19/MyPracticeBankingApp/internal/core/domain"
>>>>>>> 27c2bab4ae9973b95478eedee3812fd6c4b17ae0

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
