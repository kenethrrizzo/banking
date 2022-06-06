package utils

import (
	"encoding/json"
	"net/http"

	"github.com/kenethrrizzo/banking/logger"
)

func WriteResponse(rw http.ResponseWriter, code int, data interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		logger.Error(err.Error())
	}
}