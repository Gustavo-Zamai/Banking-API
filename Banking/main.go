package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Routes
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})
	//Start server
	http.ListenAndServe("localhost:8080", nil)

}
