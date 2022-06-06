package handlers

// Adapter: REST Handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kenethrrizzo/banking/logger"
	"github.com/kenethrrizzo/banking/service"
	util "github.com/kenethrrizzo/banking/utils"
)

type CustomerHandler struct {
	Service service.CustomerService
}

func (ch *CustomerHandler) GetAllCustomers(rw http.ResponseWriter, r *http.Request) {
	logger.Debug("getAllCustomers handler [GET]")
	var status string
	statusQuery := r.URL.Query()["status"]

	if len(statusQuery) > 0 {
		status = statusQuery[0]
	}
	customers, err := ch.Service.GetAllCustomers(status)
	if err != nil {
		util.WriteResponse(rw, err.Code, err.AssMessage())
	} else {
		util.WriteResponse(rw, http.StatusOK, customers)
	}
}

func (ch *CustomerHandler) GetCustomer(rw http.ResponseWriter, r *http.Request) {
	logger.Debug("getCustomer handler [GET]")
	vars := mux.Vars(r)
	id := vars["customer-id"]

	customer, err := ch.Service.GetCustomer(id)
	if err != nil {
		util.WriteResponse(rw, err.Code, err.AssMessage())
	} else {
		util.WriteResponse(rw, http.StatusOK, customer)
	}
}
