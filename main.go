package main

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

	log.Fatal(http.ListenAndServe(":8080", router))
}
