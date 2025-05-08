package storage

import "github.com/invocoder/task-manager/internal/types"

type Storage interface {
	CreateTask(title string, status string) (int64, error)
	GetTasksByStatus(status string, limit, offset int) ([]types.Task, error)
	UpdateTask(id int64, title string, status string) error
}
