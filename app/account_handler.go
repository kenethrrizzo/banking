package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kenethrrizzo/banking/dto"
	"github.com/kenethrrizzo/banking/logger"
	"github.com/kenethrrizzo/banking/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) newAccount(rw http.ResponseWriter, r *http.Request) {
	logger.Debug("newAccount handler [POST]")
	vars := mux.Vars(r)
	customerId := vars["customer-id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(rw, http.StatusBadRequest, err.Error())
		return
	}
	request.CustomerId = customerId
	acc, apperr := h.service.NewAccount(request)
	if apperr != nil {
		writeResponse(rw, apperr.Code, apperr.Message)
		return
	}
	writeResponse(rw, http.StatusCreated, acc)
}

func (h AccountHandler) makeTransaction(rw http.ResponseWriter, r *http.Request) {
	logger.Debug("makeTransaction handler [POST]")
	vars := mux.Vars(r)

	customerId := vars["customer-id"]
	accountId := vars["account-id"]

	var request dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(rw, http.StatusBadRequest, err.Error())
		return
	}
	request.CustomerId = customerId
	request.AccountId = accountId

	transaction, apperr := h.service.MakeTransaction(request)
	if apperr != nil {
		writeResponse(rw, apperr.Code, apperr.Message)
		return
	}
	writeResponse(rw, http.StatusCreated, transaction)
}