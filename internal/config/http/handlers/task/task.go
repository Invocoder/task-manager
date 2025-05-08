package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/invocoder/task-manager/internal/storage"
	"github.com/invocoder/task-manager/internal/types"
	"github.com/invocoder/task-manager/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info(("Creating a task"))
		var task types.Task

		err := json.NewDecoder(r.Body).Decode(&task)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError((fmt.Errorf("empty body"))))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		//request validation
		if err := validator.New().Struct(task); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}
		lastId, err := storage.CreateTask(
			task.Title,
			task.Status,
		)

		slog.Info("Task created successfully", slog.String("taskId", fmt.Sprint(lastId)))
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, map[string]string{"success": "OK"})
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

// func GetByid(storage storage.Storage) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		status := r.PathValue("status")
// 		slog.Info("getting a task", slog.String("status", status))

// 		task, err := storage.GetTaskByStatus(status)
// 		if err != nil {
// 			slog.Error("error geting task")
// 			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
// 			return
// 		}
// 		response.WriteJson(w, http.StatusOK, task)
// 	}
// }

func GetByStatusPaginated(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := r.URL.Query().Get("status")
		limitStr := r.URL.Query().Get("limit")
		offsetStr := r.URL.Query().Get("offset")

		limit, _ := strconv.Atoi(limitStr)
		offset, _ := strconv.Atoi(offsetStr)

		tasks, err := storage.GetTasksByStatus(status, limit, offset)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, tasks)
	}
}

func Update(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task types.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err := storage.UpdateTask(task.ID, task.Title, task.Status); err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, response.Response{Status: response.StatusOK})
	}
}
