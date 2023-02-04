package db

import (
	"database/sql"

	"github.com/Rokiis1/todo-list/models"
	_ "github.com/lib/pq"
)

var (
	conn *sql.DB
)

func Connect() error {
	var err error
	conn, err = sql.Open("postgres", "user=postgres password=XxQgC0sqYTWETWfO host=db.hmzdceormmealgxehegd.supabase.co port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		return err
	}

	if err = conn.Ping(); err != nil {
		return err
	}

	return nil
}

func Close() {
	conn.Close()
}

func AddTask(task models.Task) (int, error) {
	var id int

	err := conn.QueryRow("INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id", task.Title, task.Description).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func EditTask(id int, task models.Task) error {
	_, err := conn.Exec("UPDATE tasks SET title = $1, description = $2 WHERE id = $3", task.Title, task.Description, id)
	return err
}

func DeleteTask(id int) error {
	_, err := conn.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}

func GetTask(id int) (models.Task, error) {
	var task models.Task

	err := conn.QueryRow("SELECT id, title, description FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Title, &task.Description)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func GetTasks() ([]models.Task, error) {
	rows, err := conn.Query("SELECT id, title, description FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
