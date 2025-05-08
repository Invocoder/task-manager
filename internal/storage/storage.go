package storage

import "github.com/invocoder/task-manager/internal/types"

type Storage interface {
	CreateTask(title string, status string) (int64, error)
	GetTaskByStatus(status string) (types.Task, error)
}
