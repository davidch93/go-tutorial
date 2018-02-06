package endpoint

import (
	"encoding/json"
	"go-tutorial/rest-api/entity"
	"go-tutorial/rest-api/service"
	"net/http"

	"github.com/gorilla/mux"
)

// GetPeople is an API used to get list of person.
func GetPeople(w http.ResponseWriter, r *http.Request) {
	var people = service.GetPeople()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

// GetPerson is an API used to get person by ID.
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person = service.GetPerson(params["id"])
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

// CreatePerson is an API used to create new person.
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person entity.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	service.CreatePerson(person)
}

// DeletePerson is an API used to delete person by ID.
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service.DeletePerson(params["id"])
}
