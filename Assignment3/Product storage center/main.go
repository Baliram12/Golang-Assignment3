package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID          string `json:"id,omitempty"`
	Productname string `json:"Product_Name:,omitempty"`
	Model       string `json:"Model_Number:,omitempty"`
	Desc        string `json:"Price:,omitempty"`
}

var market []Product

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Fprintf(w, "Product storage cente")
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function called: getitems()")

	json.NewEncoder(w).Encode(market)
}

func GetProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range market {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

func GetmarketEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(market)
}

func CreateProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var product Product
	_ = json.NewDecoder(req.Body).Decode(&product)
	product.ID = params["id"]
	market = append(market, product)
	json.NewEncoder(w).Encode(market)
}

func DeleteProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range market {
		if item.ID == params["id"] {
			market = append(market[:index], market[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(market)
}

func main() {
	router := mux.NewRouter()
	market = append(market, Product{
		ID:          "1",
		Productname: "Nokia",
		Model:       "C3",
		Desc:        "Good Phone"})
	market = append(market, Product{
		ID:          "2",
		Productname: "Realme",
		Model:       "X7"})
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/market", GetmarketEndpoint).Methods("GET")
	router.HandleFunc("/market/{id}", GetmarketEndpoint).Methods("GET")
	router.HandleFunc("/market/{id}", CreateProductEndpoint).Methods("POST")
	router.HandleFunc("/market/{id}", DeleteProductEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}
