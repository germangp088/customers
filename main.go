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

	customer "github.com/germangp088/customers/routes"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	customer.Init()
	router.HandleFunc("/customer", customer.GetCustomers).Methods("GET")
	router.HandleFunc("/customer/{id}", customer.GetCustomer).Methods("GET")
	router.HandleFunc("/customer", customer.PostCustomer).Methods("POST")
	router.HandleFunc("/customer/{id}", customer.PutCustomer).Methods("PUT")
	router.HandleFunc("/customer/{id}", customer.DeleteCustomer).Methods("DELETE")

	routerV1 := router.PathPrefix("/v1").Subrouter()
	routerV1.HandleFunc("/customer", customer.GetCustomersV1).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
