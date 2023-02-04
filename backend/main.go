package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// API routes here

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
