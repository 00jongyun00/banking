package app

import (
	"bangking/dto"
	"bangking/errs"
	"bangking/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var router *mux.Router
var ch CustomerHandlers
var mockService *service.MockCustomerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)

	ch = CustomerHandlers{mockService}
	router = mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomer)

	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_Shoud_return_customers_with_status_code_200(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	dummyCustomers := []dto.CustomerResponse{
		{"1001", "Ashish", "New Delhi", "110011", "2000-01-01", "1"},
		{"1002", "Rob", "New Delhi", "110011", "2000-01-01", "1"},
		{"1003", "jongyun", "New Delhi", "110011", "2000-01-01", "1"},
	}
	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recoder := httptest.NewRecorder()
	router.ServeHTTP(recoder, request)

	// Assert
	assert.Equal(t, http.StatusOK, recoder.Code, "they shoud be equal")
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	mockService.EXPECT().GetAllCustomer("").Return(nil, errs.NewUnexpectedError("some database error"))
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recoder := httptest.NewRecorder()
	router.ServeHTTP(recoder, request)

	assert.Equal(t, http.StatusInternalServerError, recoder.Code, "Success while testing the status code")
}
