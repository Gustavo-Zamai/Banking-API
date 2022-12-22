package app

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Start() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()
	//Routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet) 
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomers).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomerById).Methods(http.MethodGet)
	

	//Start server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
