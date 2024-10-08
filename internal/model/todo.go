package model

import "time"

// Todo is the model for the todo endpoint.
type Todo struct {
	ID        int `gorm:"primaryKey"`
	Task      string
	Status    Status
	Priority  int
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// NewTodo returns a new instance of the todo model.
func NewTodo(task string, priority int) *Todo {
	return &Todo{
		Task:     task,
		Priority: priority,
		Status:   Created,
	}
}

// NewUpdateTodo returns a new instance of the todo model for updating.
func NewUpdateTodo(id int, task string, status Status) *Todo {
	return &Todo{
		ID:     id,
		Task:   task,
		Status: status,
	}
}

// Status is the status of the task.
type Status string

const (
	// Created is the status for a created task.
	Created = Status("created")
	// Processing is the status for a processing task.
	Processing = Status("processing")
	// Done is the status for a done task.
	Done = Status("done")
)

// Priority is the priority of the task
type Priority string

const (
	// Medium is the priority for a medium priority task.
	Medium = Priority("medium")
	// Normal is the priority of a normal priority task.
	Normal = Priority("normal")
	// High is the priority of a high priority task.
	High = Priority("high")
)

// GetPriorityInt is the function that return int corresponding of priority string.
func GetPriorityInt(priority string) int {
	switch Priority(priority) {
	case Medium:
		return 2
	case Normal:
		return 1
	case High:
		return 3

	}
	return -1
}

// IsValidStatus to validate the status
func IsValidStatus(status string) bool {
	switch Status(status) {
	case Created, Processing, Done:
		return true
	}
	return false
}

// StatusMap is a map of task status.
var StatusMap = map[Status]bool{
	Created:    true,
	Processing: true,
	Done:       true,
}
