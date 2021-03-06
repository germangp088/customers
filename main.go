package main

// Package classification Consumers API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost:8080
//     Version: 0.0.1
//     Contact: German Dario<germangp088@gmail.com> http://google.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta

import (
	"log"
	"net/http"

	customer "customers/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	Initialize()
}

//Initialize init of services.
func Initialize() {
	router := mux.NewRouter()

	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui/")))
	router.PathPrefix("/swaggerui/").Handler(sh)

	customer.Init()
	router.HandleFunc("/customer", customer.GetCustomers).Methods("GET")
	router.HandleFunc("/customer/{id}", customer.GetCustomer).Methods("GET")
	router.HandleFunc("/customer", customer.PostCustomer).Methods("POST")
	router.HandleFunc("/customer/{id}", customer.PutCustomer).Methods("PUT")
	router.HandleFunc("/customer/{id}", customer.DeleteCustomer).Methods("DELETE")

	routerV1 := router.PathPrefix("/v1").Subrouter()
	routerV1.HandleFunc("/customer", customer.GetCustomersV1).Methods("GET")

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
