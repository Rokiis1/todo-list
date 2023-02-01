package main

import (
	"fmt"
	"net/http"

	"github.com/Rokiis1/todo-list/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", handler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handler.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")

	fmt.Println("Starting server on port 3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println(err)
	}
}