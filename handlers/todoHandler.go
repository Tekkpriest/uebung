package handlers

import (
	"encoding/json"
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
		http.Error(w, "Tasks not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, "error encoding json", http.StatusInternalServerError)
		return
	}
}
