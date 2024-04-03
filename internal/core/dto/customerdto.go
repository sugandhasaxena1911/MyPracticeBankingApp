package dto

type CustomerDTO struct {
	Custid        string `json:"id"`
	Custname      string `json:"name"`
	Custbirthdate string `json:"birthDate"`
	Custcity      string `json:"city"`
	Custzipcode   string `json:"zipcode"`
	Custstatus    string `json:"status"`
}
