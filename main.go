package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Order struct (Model)
type Order struct {
	ID              string    `json:"ID"`
	createTime      string    `json:"createTime"`
	orderName       string    `json:"orderName"`
	orderitemID     int       `json:"orderitemID"`
	pricePerUnit    float32   `json:"pricePerUnit"`
	orderQuantity   int       `json:"orderQuantity"`
	deliverQuantity int       `json:"deliverQuantity"`
	product         string    `json:"product"`
	Customer        *Customer `json:"customer"`
}

// Customer struct
type Customer struct {
	userID      string `json:"userID"`
	login       string `json:"login"`
	password    string `json:"password"`
	name        string `json:"name"`
	companyID   int    `json:"companyID"`
	companyName string `json:"companyName"`
	creditCards string `json:"creditCards"`
}

// Init orders var as a slice Order struct
var orders []Order

// Get all orders
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// Get single order by orderID
func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through orders and find one with the id from the params
	for _, item := range orders {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Order{})
}

// Add new order
func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	order.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

// Update order
func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range orders {
		if item.ID == params["id"] {
			orders = append(orders[:index], orders[index+1:]...)
			var order Order
			_ = json.NewDecoder(r.Body).Decode(&order)
			order.ID = params["id"]
			orders = append(orders, order)
			json.NewEncoder(w).Encode(order)
			return
		}
	}
}

// Delete order
func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range orders {
		if item.ID == params["id"] {
			orders = append(orders[:index], orders[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(orders)
}

// Main function
func main() {
	fmt.Println("Connected to MongoDB!")
	fmt.Println("Listen to localhost:8000!")
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	orders = append(orders, Order{ID: "1", createTime: "2006-01-02T15:04:05-0700", orderName: "PO #001-I", orderitemID: 1, pricePerUnit: 1.3454, orderQuantity: 10, deliverQuantity: 5, product: "Corrugated Box", Customer: &Customer{userID: "ivan", login: "ivan", password: "12345", name: "Ivan Ivanovich", companyID: 1, companyName: "Roga & Kopyta", creditCards: "[*****-1234 *****-5678]"}})
	orders = append(orders, Order{ID: "2", createTime: "2006-01-02T15:04:05-0700", orderName: "PO #001-I", orderitemID: 1, pricePerUnit: 1.3454, orderQuantity: 10, deliverQuantity: 5, product: "Corrugated Box", Customer: &Customer{userID: "ivan", login: "ivan", password: "12345", name: "Ivan Ivanovich", companyID: 1, companyName: "Roga & Kopyta", creditCards: "[*****-1234 *****-5678]"}})

	// Route handles & endpoints
	r.HandleFunc("/orders", getOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", getOrder).Methods("GET")
	r.HandleFunc("/orders", createOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", updateOrder).Methods("PUT")
	r.HandleFunc("/orders/{id}", deleteOrder).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))

}
