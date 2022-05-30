package app

// Adapter: REST Handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kenethrrizzo/banking/logger"
	"github.com/kenethrrizzo/banking/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	logger.Debug("getAllCustomers handler [GET]")
	var status string
	statusQuery := r.URL.Query()["status"]

	if len(statusQuery) > 0 {
		status = statusQuery[0]
	}
	customers, err := ch.service.GetAllCustomers(status)
	if err != nil {
		writeResponse(rw, err.Code, err.AssMessage())
	} else {
		writeResponse(rw, http.StatusOK, customers)
	}
}

func (ch *CustomerHandler) getCustomer(rw http.ResponseWriter, r *http.Request) {
	logger.Debug("getCustomer handler [GET]")
	vars := mux.Vars(r)
	id := vars["customer-id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(rw, err.Code, err.AssMessage())
	} else {
		writeResponse(rw, http.StatusOK, customer)
	}
}

func writeResponse(rw http.ResponseWriter, code int, data interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		logger.Error(err.Error())
	}
}
