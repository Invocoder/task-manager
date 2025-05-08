package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/invocoder/task-manager/internal/config"
	"github.com/invocoder/task-manager/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	status TEXT
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateTask(title string, status string) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO tasks (title, status) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(title, status)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func (s *Sqlite) GetTaskByStatus(status string) (types.Task, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM tasks where status = ? LIMIT 1")
	if err != nil {
		return types.Task{}, err
	}

	defer stmt.Close()

	var task types.Task

	err = stmt.QueryRow(status).Scan(&task.ID, &task.Title, &task.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Task{}, fmt.Errorf("no task found with %q status", status)
		}
		return types.Task{}, fmt.Errorf("query error: %w", err)
	}
	return task, nil
}
