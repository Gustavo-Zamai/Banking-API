package app

import (
	"log"
	"net/http"
	"github.com/Gustavo-Zamai/Banking-API/domain"
	"github.com/Gustavo-Zamai/Banking-API/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	
	//Wiring
	//handlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	handlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	
	
	//Routes
	router.HandleFunc("/customers", handlers.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", handlers.getCustomer).Methods(http.MethodGet)

	//Start server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
