package sqlite

import (
	"database/sql"

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

func (s *Sqlite) GetTasksByStatus(status string, limit, offset int) ([]types.Task, error) {
	rows, err := s.Db.Query("SELECT id, title, status FROM tasks WHERE status = ? LIMIT ? OFFSET ?", status, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []types.Task
	for rows.Next() {
		var task types.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *Sqlite) UpdateTask(id int64, title string, status string) error {
	_, err := s.Db.Exec("UPDATE tasks SET title = ?, status = ? WHERE id = ?", title, status, id)
	return err
}
