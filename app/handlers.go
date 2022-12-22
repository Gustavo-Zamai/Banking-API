package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json: "name" 	 xml:"name"`
	City    string `json: "city" 	 xml:"city"`
	Zipcode string `json: "zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Gus", City: "Mococa", Zipcode: "13730000"},
		{Name: "Joaozinho", City: "Campinas", Zipcode: "13735496"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func getCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"] )
}

func createCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post Request Received")
}