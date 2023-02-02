package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Rokiis1/todo-list/database"
	"github.com/Rokiis1/todo-list/model"

	"github.com/gorilla/mux"
)

// CreateTask creates a new task in the database
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := "INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id"
	err = db.QueryRow(query, task.Title, task.Description, task.Status).Scan(&task.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// GetTasks retrieves all tasks from the database
func GetTasks(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	tasks := model.Tasks{}
	for rows.Next() {
		task := model.Task{}
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GetTask retrieves a single task from the database
func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	db, err := database.Connect()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	var task model.Task
	err = db.QueryRow("SELECT id, title, description, done FROM tasks WHERE id=$1", taskID).Scan(&task.ID, &task.Title, &task.Description, &task.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "Task not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, task)
}

func respondWithJSON(w http.ResponseWriter, i int, task model.Task) {
	panic("unimplemented")
}

func respondWithError(w http.ResponseWriter, i int, s string) {
	panic("unimplemented")
}

// UpdateTask updates a single task in the database
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	var task model.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := `UPDATE tasks SET name=?, description=?, completed=? WHERE id=?`
	_, err = db.Exec(query, task.Name, task.Description, task.Completed, taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// DeleteTask deletes a single task from the database
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := `DELETE FROM tasks WHERE id=?`
	_, err = db.Exec(query, taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}