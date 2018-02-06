package main

import (
	"go-tutorial/rest-api/endpoint"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", endpoint.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", endpoint.GetPerson).Methods("GET")
	router.HandleFunc("/people", endpoint.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", endpoint.DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
