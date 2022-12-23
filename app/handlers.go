package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"github.com/Gustavo-Zamai/Banking-API/service"
	//"fmt"
	//"github.com/gorilla/mux"
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
	//customers := []Customer{
	//{Name: "Gus", City: "Mococa", Zipcode: "13730000"},
	//	{Name: "Joaozinho", City: "Campinas", Zipcode: "13735496"},
	//}

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
