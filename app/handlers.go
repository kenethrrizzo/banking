package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/kenethrrizzo/banking/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"FullName"`
	City    string `json:"city" xml:"City"`
	Zipcode string `json:"zip_code" xml:"ZipCode"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(rw).Encode(customers)
	} else {
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(customers)
	}
}
