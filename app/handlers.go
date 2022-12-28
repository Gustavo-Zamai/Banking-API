package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"github.com/Gustavo-Zamai/Banking-API/service"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json: "name" 	 xml:"name"`
	City    string `json: "city" 	 xml:"city"`
	Zipcode string `json: "zip_code" xml:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["customer_id"]

    customer, err := ch.service.GetCustomer(id)

    if err!= nil {
        w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
    }else{
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
