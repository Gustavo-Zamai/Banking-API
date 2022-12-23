package app

import (
	"log"
	"net/http"

	"github.com/Gustavo-Zamai/Banking-API/domain"
	"github.com/Gustavo-Zamai/Banking-API/service"
	"github.com/gorilla/mux"
)

func Start() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()
	
	//Wiring
	handlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	//Routes
	router.HandleFunc("/customers", handlers.getAllCustomers).Methods(http.MethodGet)

	

	//Start server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
