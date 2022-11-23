package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Customer struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
	Contacted bool   `json:"contacted"`
}

type RequestCustomer struct {
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customers = make(map[string]Customer)

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if _, ok := customers[id]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers[id])
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(customers)
	}

}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newCustomer RequestCustomer

	err := json.NewDecoder(r.Body).Decode(&newCustomer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t := time.Now()
	customerInsertId := t.Format("20060102150405")
	customers[customerInsertId] = Customer{
		Id:        customerInsertId,
		Name:      newCustomer.Name,
		Role:      newCustomer.Role,
		Email:     newCustomer.Email,
		Phone:     newCustomer.Phone,
		Contacted: newCustomer.Contacted}

	if _, ok := customers[customerInsertId]; ok {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(customers[customerInsertId])
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(customers)
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	var updateCustomer RequestCustomer

	err := json.NewDecoder(r.Body).Decode(&updateCustomer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	delete(customers, id)
	customers[id] = Customer{
		Id:        id,
		Name:      updateCustomer.Name,
		Role:      updateCustomer.Role,
		Email:     updateCustomer.Email,
		Phone:     updateCustomer.Phone,
		Contacted: updateCustomer.Contacted}

	if _, ok := customers[id]; ok {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(customers[id])
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(customers)
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if _, ok := customers[id]; ok {
		delete(customers, id)
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(customers)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(customers)
	}
}

func showInfoPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func main() {
	// TODO
	customersData := []Customer{{Id: "1", Name: "Ono", Role: "Ono role", Email: "Ono mail", Phone: 1111111, Contacted: false},
		{Id: "2", Name: "Dus", Role: "Dus role", Email: "Dus email", Phone: 2222222, Contacted: false},
		{Id: "3", Name: "Tros", Role: "Tros role", Email: "Tros email", Phone: 333333, Contacted: false},
		{Id: "4", Name: "Cotro", Role: "Cotro role", Email: "Cotro email", Phone: 4444444, Contacted: false},
		{Id: "5", Name: "Conco", Role: "Conco role", Email: "Conco email", Phone: 55555555, Contacted: false}}

	for _, customerJson := range customersData {
		customers[customerJson.Id] = Customer{
			Id:        customerJson.Id,
			Name:      customerJson.Name,
			Role:      customerJson.Role,
			Email:     customerJson.Email,
			Phone:     customerJson.Phone,
			Contacted: customerJson.Contacted}
	}

	jsonData, err := json.Marshal(customers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))

	router := mux.NewRouter()
	http.Handle("/", http.FileServer(http.Dir("./static")))
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PATCH")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	router.HandleFunc("/", showInfoPage).Methods("GET")

	fmt.Println("Server is starting on port 3000...")
	// Pass the customer router into ListenAndServe
	http.ListenAndServe(":3000", router)

}
