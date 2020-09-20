package main

import (
	"go-crud-api/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/events", services.CreateEvent).Methods("POST")
	router.HandleFunc("/events", services.GetEvents).Methods("GET")
	router.HandleFunc("/events/{id}", services.GetAnEvent).Methods("GET")
	router.HandleFunc("/events/{id}", services.UpdateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", services.DeleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
