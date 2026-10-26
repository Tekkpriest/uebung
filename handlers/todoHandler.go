package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/uebung/data"
	"github.com/uebung/models"
)

type ToDoSource interface {
	GetTasks() ([]models.Task, error)
	GetTaskByID(id int) (models.Task, error)
}

type ToDoHandler struct {
	ToDos ToDoSource
}

func NewToDoHandler(t ToDoSource) *ToDoHandler {
	return &ToDoHandler{ToDos: t}
}

func (t *ToDoHandler) GetToDos(w http.ResponseWriter, r *http.Request) {
	todos, err := t.ToDos.GetTasks()
	if err != nil {
		http.Error(w, "Tasks could not be loaded", http.StatusInternalServerError)
		slog.Error("error loading todos", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(todos); err != nil {
		slog.Error("could not encode todos", "error", err)
	}
}

func (t *ToDoHandler) GetToDoByID(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		slog.Error("error loading todo id", "error", err)
		return
	}

	todo, err := t.ToDos.GetTaskByID(id)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			http.Error(w, "todo not found", http.StatusNotFound)
			return
		}

		http.Error(w, "todo could not be loaded", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		slog.Error("could not encode todo", "error", err)
	}
}
