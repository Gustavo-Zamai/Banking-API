package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string "json "
	City    string
	Zipcode string
}

func main() {
	//Routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	//Start server
	http.ListenAndServe("localhost:8080", nil)

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Gus", City: "Mococa", Zipcode: "13730000"},
		{Name: "Joaozinho", City: "Campinas", Zipcode: "13735496"},
	}
	json.NewEncoder(w).Encode(customers)
}
