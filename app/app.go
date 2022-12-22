package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()

	//Routes
	mux.HandleFunc("/greet", greet) 
	mux.HandleFunc("/customers", getAllCustomers)

	//Start server
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
