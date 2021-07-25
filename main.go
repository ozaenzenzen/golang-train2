package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Car struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type Cars []Car

func allCars(w http.ResponseWriter, r *http.Request) {
	car := Cars{
		Car{Title: "Test Titile", Desc: "Test Desc", Content: "Test Content"},
	}
	fmt.Println("Endpoint hit: All Cars Endpoint")
	json.NewEncoder(w).Encode(car)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/cars", allCars)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
