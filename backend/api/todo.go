package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Rokiis1/todo-list/db"
	"github.com/Rokiis1/todo-list/errors"
	"github.com/Rokiis1/todo-list/models"
	"github.com/gorilla/mux"
)

// AddTaskHandler creates a new task in the database.
func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		errors.WriteError(err, w)
		return
	}

	newID, err := db.AddTask(task)
	if err != nil {
		errors.WriteError(err, w)
		return
	}

	task.ID = newID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// EditTaskHandler updates a task in the database.
func EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskIDStr := vars["id"]

	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		errors.WriteError(err, w)
		return
	}

	var task models.Task

	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		errors.WriteError(err, w)
		return
	}

	err = db.EditTask(taskID, task)
	if err != nil {
		errors.WriteError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// DeleteTaskHandler deletes a task from the database.
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskIDStr := vars["id"]

	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		errors.WriteError(err, w)
		return
	}

	deleteErr := db.DeleteTask(taskID)
	if err != nil {
		errors.WriteError(deleteErr, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.GetTasks()
	if err != nil {
		errors.WriteError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
