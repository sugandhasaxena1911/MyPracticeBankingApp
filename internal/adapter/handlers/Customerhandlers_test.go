package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/core/dto"
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/core/ports/mocks/service"
)

func Test_GetAllCustomers_return_custs_200OK(t *testing.T) {
	//arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockCustomerService(ctrl)
	ch := Custhandlers{mockService}
	dummycusts := []dto.CustomerDTO{
		{Custid: "1", Custname: "Sugandha", Custbirthdate: "2023-11-05", Custcity: "Lucknow", Custzipcode: "6766", Custstatus: "active"},
		{Custid: "2", Custname: "dummy", Custbirthdate: "2023-11-05", Custcity: "Lucknow", Custzipcode: "67696", Custstatus: "active"},
	}

	mockService.EXPECT().FindAllCustomers().Return(dummycusts, nil)
	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.GetAllCustomers)
	req, err := http.NewRequest(http.MethodGet, "/customers", nil)
	if err != nil {
		log.Fatalln("request error ")
	}

	recorder := httptest.NewRecorder() // recorder is response
	router.ServeHTTP(recorder, req)
	//assert
	if recorder.Code != http.StatusOK {
		t.Error("failed while testing GEt Customers , 200 OK not recieved ")

	}

}

func Test_GetAllCustomers_return_500_internal_serverError(t *testing.T) {

	//arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockCustomerService(ctrl)
	ch := Custhandlers{mockService}

	mockService.EXPECT().FindAllCustomers().Return(nil, error.NewInternalServerAppError("Internal DB Error"))

	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.GetAllCustomers)
	req, err := http.NewRequest(http.MethodGet, "/customers", nil)
	if err != nil {
		log.Fatalln("request error ")
	}
	recorder := httptest.NewRecorder() // recorder is response
	router.ServeHTTP(recorder, req)
	//assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("failed while testing GEt Customers , 500 internal server not recieved ")

	}
}
