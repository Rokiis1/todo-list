package main

import (
	"log"
	"net/http"

	"github.com/Rokiis1/todo-list/api"
	"github.com/Rokiis1/todo-list/db"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()

	err := db.Connect()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Define the API routes
	router.HandleFunc("/tasks", api.GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks", api.AddTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", api.EditTaskHandler).Methods("PUT")
	router.HandleFunc("/tasks/{id}", api.DeleteTaskHandler).Methods("DELETE")

	// Start the API server
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}

}
