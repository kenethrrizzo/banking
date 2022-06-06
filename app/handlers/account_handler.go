package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kenethrrizzo/banking/dto"
	"github.com/kenethrrizzo/banking/logger"
	"github.com/kenethrrizzo/banking/service"
	util "github.com/kenethrrizzo/banking/utils"
)

type AccountHandler struct {
	Service service.AccountService
}

func (h AccountHandler) NewAccount(rw http.ResponseWriter, r *http.Request) {
	logger.Debug("newAccount handler [POST]")
	vars := mux.Vars(r)
	customerId := vars["customer-id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		util.WriteResponse(rw, http.StatusBadRequest, err.Error())
		return
	}
	request.CustomerId = customerId
	acc, apperr := h.Service.NewAccount(request)
	if apperr != nil {
		util.WriteResponse(rw, apperr.Code, apperr.Message)
		return
	}
	util.WriteResponse(rw, http.StatusCreated, acc)
}

func (h AccountHandler) MakeTransaction(rw http.ResponseWriter, r *http.Request) {
	logger.Debug("makeTransaction handler [POST]")
	vars := mux.Vars(r)

	customerId := vars["customer-id"]
	accountId := vars["account-id"]

	var request dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		util.WriteResponse(rw, http.StatusBadRequest, err.Error())
		return
	}
	request.CustomerId = customerId
	request.AccountId = accountId

	transaction, apperr := h.Service.MakeTransaction(request)
	if apperr != nil {
		util.WriteResponse(rw, apperr.Code, apperr.Message)
		return
	}
	util.WriteResponse(rw, http.StatusCreated, transaction)
}