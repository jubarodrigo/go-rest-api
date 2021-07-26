package controller

import (
	"encoding/json"
	"net/http"
	"github.com/jubarodrigo/api-rest/model"
	"github.com/jubarodrigo/api-rest/service"
)

var people []model.Person

func GetPeople(w http.ResponseWriter, r *http.Request) {

	people = append(people, model.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &model.Address{City: "City X", State: "State X"}})
	people = append(people, model.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &model.Address{City: "City Z", State: "State Y"}})
	people = append(people, model.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {}

func CreatePerson(w http.ResponseWriter, request *http.Request) {

	service.CreatePerson(request)

}

func DeletePerson(w http.ResponseWriter, r *http.Request) {}