//go:generate swagger generate spec
package customer

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Type of customers
// swagger:model
type Customer struct {
	// the id for this customer
	//
	// required: true
	// min: 1
	ID int `json:"id,omitempty"`
	// the name for this customer
	// required: true
	Firstname string `json:"firstname,omitempty"`
	// the lastname for this customer
	// required: true
	Lastname string `json:"lastname,omitempty"`
}

var customers []Customer

//Init initialice customers
func Init() {
	customers = append(customers, Customer{ID: 1, Firstname: "John", Lastname: "Doe"})
	customers = append(customers, Customer{ID: 2, Firstname: "Mark", Lastname: "Antony"})
	customers = append(customers, Customer{ID: 3, Firstname: "Sara", Lastname: "Flowers"})
	customers = append(customers, Customer{ID: 4, Firstname: "Sherman", Lastname: "Dario"})
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
	//
	//     Responses:
	//       200: []Customer
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

//PutCustomer Update a customer
func PutCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Invalid param id")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var customer Customer
	_ = json.NewDecoder(r.Body).Decode(&customer)

	for index, item := range customers {
		if item.ID == i {

			item.Firstname = customer.Firstname
			item.Lastname = customer.Lastname

			customers[index] = item

			json.NewEncoder(w).Encode(customers)
		}
	}
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
