package storage

type Storage interface {
	CreateTask(title string, status string) (int64, error)
}