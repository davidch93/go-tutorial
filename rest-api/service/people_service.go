package service

import (
	"github.com/davidch93/go-tutorial/rest-api/entity"
)

var people []entity.Person

func init() {
	people = append(people, entity.Person{ID: "1", Firstname: "John", Lastname: "Doe",
		Address: &entity.Address{City: "City X", State: "State X"}})
	people = append(people, entity.Person{ID: "2", Firstname: "Koko", Lastname: "Doe",
		Address: &entity.Address{City: "City Z", State: "State Y"}})
	people = append(people, entity.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
}

// GetPeople is a service used to get all person.
func GetPeople() []entity.Person {
	return people
}

// GetPerson is a service used to get person by ID.
func GetPerson(ID string) entity.Person {
	var defaultValue entity.Person
	for _, person := range people {
		if person.ID == ID {
			return person
		}
	}
	return defaultValue
}

// CreatePerson is a service used to create new person.
func CreatePerson(person entity.Person) {
	people = append(people, person)
}

// DeletePerson is a service used to delete person by ID.
func DeletePerson(ID string) {
	for index, item := range people {
		if item.ID == ID {
			people[index] = people[len(people)-1]
			people = people[:len(people)-1]
			break
		}
	}
}
