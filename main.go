package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Type os customers
type Customer struct {
	ID        int    `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

var customers []Customer

func main() {
	customers = append(customers, Customer{ID: 1, Firstname: "John", Lastname: "Doe"})
	customers = append(customers, Customer{ID: 2, Firstname: "Mark", Lastname: "Antony"})
	customers = append(customers, Customer{ID: 3, Firstname: "Sara", Lastname: "Flowers"})
	customers = append(customers, Customer{ID: 4, Firstname: "Sherman", Lastname: "Dario"})

	router := mux.NewRouter()
	router.HandleFunc("/customer", GetCustomers).Methods("GET")
	router.HandleFunc("/customer/{id}", GetCustomer).Methods("GET")
	router.HandleFunc("/customer", PostCustomer).Methods("POST")
	router.HandleFunc("/customer/{id}", DeleteCustomer).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

//GetCustomers Get all customers
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(customers)
}

//GetCustomer Get a customer by id
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Invalid param id")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	for _, item := range customers {
		if item.ID == i {
			json.NewEncoder(w).Encode(item)
		}
	}
}

//PostCustomer Add a new customer
func PostCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	_ = json.NewDecoder(r.Body).Decode(&customer)
	id := (len(customers) + 1)
	customer.ID = id
	customers = append(customers, customer)
	json.NewEncoder(w).Encode(customer.ID)
}

//DeleteCustomer Delete a customer by id
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Invalid param id")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var newcustomers []Customer

	// redneck solution
	for index, item := range customers {
		if item.ID != i {
			newcustomers = append(newcustomers, customers[index])
		}
	}

	customers = newcustomers

	json.NewEncoder(w).Encode(customers)
}
