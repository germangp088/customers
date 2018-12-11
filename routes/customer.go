//go:generate swagger generate spec
package customer

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// HTTP status code 200 and repository model in data
// swagger:response repoResp
type swaggRepoResp struct {
	// in:body
	Body struct {
		// HTTP status code 200/201
		Code int `json:"code"`
		// Repository model
		Data int `json:"id"`
	}
}

// HTTP status code 500/400
// swagger:response error
type swaggRepoRespError struct {
	// in:body
	Body struct {
		// HTTP status code 200/201
		Code int `json:"code"`
		// Repository model
		Data string `json:"error"`
	}
}

//Customer Struct of customers
// swagger:model
type Customer struct {
	// The id for this customer
	//
	// required: true
	// min: 1
	ID int `json:"id,omitempty"`
	// The name for this customer
	// required: true
	Firstname string `json:"firstname,omitempty"`
	// The lastname for this customer
	// required: true
	Lastname string `json:"lastname,omitempty"`
}

// Customers Customers contains information about customers available for public API
// swagger:model Customers
type Customers []*Customer

var customers []Customer

//Init initialice customers
func Init() {
	customers = append(customers, Customer{ID: 1, Firstname: "John", Lastname: "Doe"})
	customers = append(customers, Customer{ID: 2, Firstname: "Mark", Lastname: "Antony"})
	customers = append(customers, Customer{ID: 3, Firstname: "Sara", Lastname: "Flowers"})
	customers = append(customers, Customer{ID: 4, Firstname: "Sherman", Lastname: "Dario"})
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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
	json.NewEncoder(w).Encode(customers)
}

func GetCustomersV1(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(customers)
}

//GetCustomer Get a customer by id
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /customer/{id} customer Customer
	// ---
	// summary: Specific customer by id.
	// description: If id is not valid, Error Bad Request (400) will be returned.
	// parameters:
	// - name: id
	//   in: /{id}
	//   description: customer id
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/definitions/Customer"
	//   "400":
	//     "$ref": "#/responses/error"
	//   "500":
	//     "$ref": "#/responses/error"
	setupResponse(&w, r)
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
	// swagger:operation POST /customer customer Customer
	// ---
	// summary: Add new customer.
	// parameters:
	// - name: Firstname
	//   in: body
	//   description: Firstname
	//   type: string
	//   required: true
	// - name: Lastname
	//   in: body
	//   description: Lastname
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/repoResp"
	//   "500":
	//     "$ref": "#/responses/error"
	setupResponse(&w, r)
	var customer Customer
	_ = json.NewDecoder(r.Body).Decode(&customer)
	id := (len(customers) + 1)
	customer.ID = id
	customers = append(customers, customer)
	json.NewEncoder(w).Encode(customer.ID)
}

//PutCustomer Update a customer
func PutCustomer(w http.ResponseWriter, r *http.Request) {
	// swagger:operation PUT /customer/{id} customer Customer
	// ---
	// summary: Update customer by id.
	// description: If id is not valid, Error Bad Request (400) will be returned.
	// parameters:
	// - name: id
	//   in: /customer/{id}
	//   description: customer id
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/definitions/Customers"
	//   "400":
	//     "$ref": "#/responses/error"
	//   "500":
	//     "$ref": "#/responses/error"
	setupResponse(&w, r)
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
	// swagger:operation DELETE /customer/{id} customer Customer
	// ---
	// summary: Delete customer by id.
	// description: If id is not valid, Error Bad Request (400) will be returned.
	// parameters:
	// - name: id
	//   in: /customer/{id}
	//   description: customer id
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/definitions/Customers"
	//   "400":
	//     "$ref": "#/responses/error"
	//   "500":
	//     "$ref": "#/responses/error"
	setupResponse(&w, r)
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
