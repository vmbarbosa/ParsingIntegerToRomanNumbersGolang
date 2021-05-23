package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Crear enrutador
	Router()
}

func Router() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/api/conversor/{id}", conversor).Methods("GET")

	http.ListenAndServe(":3000", router)

}
