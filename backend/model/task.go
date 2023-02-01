package model

// Task represents a single task in the to-do list
type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
}

// Tasks is a slice of Task
type Tasks []Task
