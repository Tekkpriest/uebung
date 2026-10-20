package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/uebung/models"
)

type ToDoSource interface {
	GetTasks() ([]models.Task, error)
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
