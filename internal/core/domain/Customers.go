package domain

import "github.com/sugandhasaxena19/MyPracticeBankingApp/internal/core/dto"

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
