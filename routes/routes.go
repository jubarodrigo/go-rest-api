package routes

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jubarodrigo/api-rest/controller"
)

func Router() {

    router := mux.NewRouter()

	router.HandleFunc("/", controller.GetPeople).Methods("GET")
	router.HandleFunc("/{id}", controller.GetPerson).Methods("GET")
	router.HandleFunc("/",controller.CreatePerson).Methods("POST")
	router.HandleFunc("/{id}", controller.DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8088", router))
	
}