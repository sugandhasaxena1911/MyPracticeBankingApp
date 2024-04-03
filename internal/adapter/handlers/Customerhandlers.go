package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sugandhasaxena19/MyPracticeBankingApp/helpers/error"
	"github.com/sugandhasaxena19/MyPracticeBankingApp/internal/core/dto"
	"github.com/sugandhasaxena19/MyPracticeBankingApp/internal/core/ports/service"
)

type Custhandlers struct {
	Custservice service.CustomerService
}

func (c *Custhandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	// get the status from the url
	urlvalues := r.URL.Query()
	status := urlvalues.Get("status")
	log.Println("Status is ", status)
	var custs []dto.CustomerDTO
	var err *error.AppError
	if status != "" {
		custs, err = c.Custservice.FindAllCustomersByStatus(status)

	} else {
		custs, err = c.Custservice.FindAllCustomers()

	}

	if err != nil {
		writeResponse(w, err.Code, err.Getmessage())
		return
	}
	writeResponse(w, http.StatusOK, custs)
}

func (c *Custhandlers) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	var custid string
	if v, ok := m["id"]; ok {
		custid = v
	}
	log.Println("Id fetched from request is ", custid)
	customer, err := c.Custservice.FindCustomerByID(custid)
	if err != nil {
		writeResponse(w, err.Code, err.Getmessage())
		return
	}
	writeResponse(w, http.StatusOK, customer)
}

func (c *Custhandlers) PostCustomer(w http.ResponseWriter, r *http.Request) {
	dtocust := dto.CustomerDTO{}
	er := json.NewDecoder(r.Body).Decode(&dtocust)
	if er != nil {
		e := error.NewBadRequestAppError("Invalid request")
		writeResponse(w, http.StatusBadRequest, e.Getmessage())
	}
	log.Println(dtocust)
	dtocust, err := c.Custservice.CreateCustomer(dtocust)
	if err != nil {
		writeResponse(w, err.Code, err.Getmessage())
		return
	}
	writeResponse(w, http.StatusCreated, dtocust)

}

func writeResponse(w http.ResponseWriter, errcode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(errcode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Panicln(err)
	}
}
