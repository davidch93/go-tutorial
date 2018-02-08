package main

import (
	"log"
	"net/http"

	"github.com/davidch93/go-tutorial/rest-api/endpoint"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", endpoint.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", endpoint.GetPerson).Methods("GET")
	router.HandleFunc("/find-people", endpoint.FindPeople).Methods("GET")
	router.HandleFunc("/people", endpoint.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", endpoint.DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
