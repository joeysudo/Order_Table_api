package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Person struct{
	user_id string `jason:"user_id,omitempty"bson:"user_id,omitempty"`
	login string `jason:"login,omitempty"bson:"login,omitempty"`
	password string `jason:"password,omitempty"bson:"password,omitempty"`
	name string `jason:"name,omitempty"bson:"name,omitempty"`
	company_id    int     `jason:"company_id,omitempty"bson:"company_id,omitempty"`
	company_name  string   `jason:"company_name,omitempty"bson:"company_name,omitempty"`
	credit_cards string `jason:"credit_cards,omitempty"bson:"credit_cards,omitempty"`
}

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("test").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) { 
	response.Header().Set("content-type", "application/json")
	var people []Person
	collection := client.Database("test").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

func main() {
	// Set client options
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// Connect to MongoDB
client, err := mongo.Connect(context.TODO(), clientOptions)
if err != nil {
    log.Fatal(err)
}
// Check the connection
err = client.Ping(context.TODO(), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Connected to MongoDB!")
collection := client.Database("test").Collection("people")
ivan:=Person{"ivan","ivan","12345","Ivan Ivanovich",1,"Roga & Kopyta","*****-1234,*****-5678"}
insertResult, err := collection.InsertOne(context.TODO(), )
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	router := mux.NewRouter()
	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	http.ListenAndServe(":8000", router)
}