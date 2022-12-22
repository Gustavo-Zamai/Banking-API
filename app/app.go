package app

import (
	"log"
	"net/http"
)

func Start() {
	//Routes
	http.HandleFunc("/greet", greet) 
	http.HandleFunc("/customers", getAllCustomers)

	//Start server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
