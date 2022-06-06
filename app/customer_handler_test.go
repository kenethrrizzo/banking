package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/kenethrrizzo/banking/dto"
	errs "github.com/kenethrrizzo/banking/error"
	"github.com/kenethrrizzo/banking/mocks/service"
)

func TestShouldReturnCustomersWithStatusCode200(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockCustomerService(ctrl)

	dummyCustomers := []dto.CustomerResponse{
		{Id: "1001", Name: "Keneth", City: "Guayaquil", ZipCode: "12345", DateOfBirth: "2000-02-12", Status: "1"},
		{Id: "1002", Name: "Camila", City: "Guayaquil", ZipCode: "12344", DateOfBirth: "2000-23-02", Status: "1"},
		{Id: "1003", Name: "Carlos", City: "Quito", ZipCode: "12343", DateOfBirth: "2000-04-08", Status: "1"},
	}

	mockService.EXPECT().GetAllCustomers("").Return(dummyCustomers, nil)

	customerHandler := CustomerHandler{mockService}

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.getAllCustomers)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func TestShouldReturnStatusCode500WithErrorMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockCustomerService(ctrl)

	mockService.EXPECT().GetAllCustomers("").Return(nil, errs.NewUnexpectedError("Some database error"))

	customerHandler := CustomerHandler{mockService}

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.getAllCustomers)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
