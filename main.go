package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Order struct (Model)
type Order struct {
	OrderID         string    `json:"order_id"`
	OrderName       string    `json:"order_name"`
	CreateTime      time.Time `json:"create_time"`
	PricePerUnit    float64   `json:"price_per_unit"`
	OrderQuantity   int64     `json:"order_quantity"`
	DeliverQuantity float64   `json:"deliver_quantity"`
	Product         string    `json:"product"`
	Customer        *Customer `json:"customer"`
}

// Customer struct
type Customer struct {
	Name        string `json:"customer_name"`
	CompanyName string `json:"company_name"`
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
		if item.OrderID == params["id"] {
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
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

// Update order
func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range orders {
		if item.OrderID == params["id"] {
			orders = append(orders[:index], orders[index+1:]...)
			var order Order
			_ = json.NewDecoder(r.Body).Decode(&order)
			order.OrderID = params["id"]
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
		if item.OrderID == params["id"] {
			orders = append(orders[:index], orders[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(orders)
}

// Main function
func main() {

	// @todo: add database
	csvFile, _ := os.Open("test_data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		record, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		t, _ := time.Parse(time.RFC3339, record[1])
		p, _ := strconv.ParseFloat(record[4], 64)
		oq, _ := strconv.ParseInt(record[5], 10, 64)
		dq, _ := strconv.ParseFloat(record[6], 64)
		orders = append(orders, Order{
			OrderID:         record[0],
			OrderName:       record[2],
			CreateTime:      t,
			PricePerUnit:    p,
			OrderQuantity:   oq,
			DeliverQuantity: dq,
			Product:         record[7],
			Customer: &Customer{
				Name:        record[11],
				CompanyName: record[13],
			},
		})
	}
	orderJSON, _ := json.Marshal(orders)
	fmt.Println(string(orderJSON))
	fmt.Println("Listen to localhost:8000!")
	// Init router
	r := mux.NewRouter()
	// Route handles & endpoints
	r.HandleFunc("/orders", getOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", getOrder).Methods("GET")
	r.HandleFunc("/orders", createOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", updateOrder).Methods("PUT")
	r.HandleFunc("/orders/{id}", deleteOrder).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))

}
