package storage

import (
	"errors"
	"todo-list/models"
)

type MemoryStore struct {
	tasks  map[int]models.Task
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

func (s *MemoryStore) AddTask(title string) models.Task {
	task := models.Task{
		ID:       s.nextID,
		Title:    title,
		Complete: false,
	}
	s.tasks[s.nextID] = task
	s.nextID++
	return task
}

func (s *MemoryStore) GetTasks() []models.Task {
	var tasks []models.Task
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (s *MemoryStore) GetTaskByID(id int) (models.Task, error) {
	task, exists := s.tasks[id]
	if !exists {
		return models.Task{}, errors.New("Task not found.")
	}
	return task, nil
}

func (s *MemoryStore) UpdateTask(id int, complete bool) error {
	task, exists := s.tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	task.Complete = complete
	s.tasks[id] = task
	return nil
}

func (s *MemoryStore) DeleteTask(id int) error {
	if _, exists := s.tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(s.tasks, id)
	return nil
}
