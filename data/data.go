package data

import (
	"errors"
	"sync"

	"github.com/uebung/models"
)

type StoreTasks struct {
	mu    sync.RWMutex
	tasks []models.Task
}

func NewStoreTasks() *StoreTasks {
	return &StoreTasks{
		tasks: []models.Task{
			{
				ID:     1,
				Title:  "Einkaufen",
				Status: false,
			},
			{
				ID:     2,
				Title:  "Musik hören",
				Status: false,
			},
			{
				ID:     3,
				Title:  "Lernen",
				Status: false,
			},
		},
	}
}

func (s *StoreTasks) GetTasks() ([]models.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.tasks, nil
}

func (s *StoreTasks) GetTaskByID(id int) (models.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return models.Task{}, errors.New("task not found")
}
