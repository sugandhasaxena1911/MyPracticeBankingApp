package service

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_validateStatus(t *testing.T) {
	//arrange
	status := "invalid"
	// act
	apperr := validateStatus(status)
	//assert
	if apperr.Message != fmt.Sprintf("invalid status :%s", status) {
		t.Error("Invalid message recieved while testing customer status")
	}
	if apperr.Code != http.StatusBadRequest {
		t.Error("Invalid http response code recieved while testing customer status ")
	}

	status = "active"
	// act
	apperr = validateStatus(status)
	//assert
	if apperr != nil {
		t.Error("Invalid error recived while testing status")
	}
}
