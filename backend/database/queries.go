package database

import "database/sql"

// createTaskInDB adds a new task to the database
func createTaskInDB(db *sql.DB, task *Task) error {
	_, err := db.Exec("INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)", task.Title, task.Description, task.Status)
	return err
}

// getTaskFromDB retrieves a single task from the database
func getTaskFromDB(db *sql.DB, taskID int) (*Task, error) {
	var task Task
	err := db.QueryRow("SELECT id, title, description, status FROM tasks WHERE id = $1", taskID).Scan(&task.ID, &task.Title, &task.Description, &task.Status)
	return &task, err
}

// updateTaskInDB updates a single task in the database
func updateTaskInDB(db *sql.DB, taskID int, task *Task) error {
	_, err := db.Exec("UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4", task.Title, task.Description, task.Status, taskID)
	return err
}

// deleteTaskInDB deletes a single task from the database
func deleteTaskInDB(db *sql.DB, taskID int) error {
    // Prepare a DELETE statement to remove a task with a specific ID from the database
    stmt, err := db.Prepare("DELETE FROM tasks WHERE id = $1")
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Execute the DELETE statement and return any errors
    _, err = stmt.Exec(taskID)
    return err
}