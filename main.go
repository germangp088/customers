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
/*package main

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
*/

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/customer", GetCustomers).Methods("GET")
	/*router.HandleFunc("/customer/{id}", customer.GetCustomer).Methods("GET")
	router.HandleFunc("/customer", customer.PostCustomer).Methods("POST")
	router.HandleFunc("/customer/{id}", customer.PutCustomer).Methods("PUT")
	router.HandleFunc("/customer/{id}", customer.DeleteCustomer).Methods("DELETE")

	routerV1 := router.PathPrefix("/v1").Subrouter()
	routerV1.HandleFunc("/customer", customer.GetCustomersV1).Methods("GET")*/

	log.Fatal(http.ListenAndServe(":8080", router))
}

//GetCustomers Get all customers
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /customer customers
	//
	// GetCustomers Get all customers
	//
	// This will show all available customers by default.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//       200: []Customer
	//  	 400: swaggRepoRespError
	//  	 500: swaggRepoRespError
	setupResponse(&w, r)

	var customers []string

	customers = append(customers, "John")
	customers = append(customers, "Mark")
	json.NewEncoder(w).Encode(customers)
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
