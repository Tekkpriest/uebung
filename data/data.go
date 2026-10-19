package data

import "github.com/uebung/models"

type StoreTasks struct {
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
	return s.tasks, nil
}
