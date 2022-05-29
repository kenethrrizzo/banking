package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kenethrrizzo/banking/dto"
	"github.com/kenethrrizzo/banking/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) newAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(rw, http.StatusBadRequest, err.Error())
		return
	}
	request.CustomerId = customerId
	acc, accerror := h.service.NewAccount(request)
	if accerror != nil {
		writeResponse(rw, accerror.Code, accerror.Message)
		return
	}
	writeResponse(rw, http.StatusCreated, acc)
}