package domain

<<<<<<< HEAD
import "github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/core/dto"
=======
import "github.com/sugandhasaxena19/MyPracticeBankingApp/internal/core/dto"
>>>>>>> 27c2bab4ae9973b95478eedee3812fd6c4b17ae0

type Customers struct {
	Custid        string
	Custname      string
	Custbirthdate string
	Custcity      string
	Custzipcode   string
	Custstatus    string
}

func (customer Customers) ToCustomerDTO() dto.CustomerDTO {

	return dto.CustomerDTO{Custid: customer.Custid,
		Custname:      customer.Custname,
		Custbirthdate: customer.Custbirthdate,
		Custcity:      customer.Custcity,
		Custzipcode:   customer.Custzipcode,
		Custstatus:    customer.GetStatusAsText()}

}
func (customer Customers) GetStatusAsText() string {
	var status string
	switch customer.Custstatus {
	case "0":
		status = "inactive"
	case "1":
		status = "active"

	}
	return status

}
