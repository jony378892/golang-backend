package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am Jony. I'm a softwear engineer")
}

type Product struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "Please send me a get request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)

	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/products", getProducts)

	fmt.Println("Server running on 3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	prod1 := Product{
		Id:          1,
		Title:       "Orange",
		Description: "This is Orange",
	}

	prod2 := Product{
		Id:          2,
		Title:       "Orange",
		Description: "This is Orange",
	}
	prod3 := Product{
		Id:          3,
		Title:       "Orange",
		Description: "This is Orange",
	}

	productList = append(productList, prod1)
	productList = append(productList, prod2)
	productList = append(productList, prod3)
}
