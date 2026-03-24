package task

import (
	"errors"
	"task-cli/domain"
	"time"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type Repository interface {
	Load() ([]domain.Task, error)
	Save(tasks []domain.Task) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Add(description string) (domain.Task, error) {
	tasks, err := s.repo.Load()
	if err != nil {
		return domain.Task{}, err
	}

	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	newTask := domain.Task{
		ID:          maxID + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	return newTask, s.repo.Save(tasks)
}

func (s *Service) Update(id int, description string) error {
	tasks, err := s.repo.Load()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return ErrTaskNotFound
	}

	return s.repo.Save(tasks)
}

func (s *Service) Delete(id int) error {
	tasks, err := s.repo.Load()
	if err != nil {
		return err
	}

	found := false
	var updatedTasks []domain.Task
	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		updatedTasks = append(updatedTasks, t)
	}

	if !found {
		return ErrTaskNotFound
	}

	if updatedTasks == nil {
		updatedTasks = []domain.Task{}
	}

	return s.repo.Save(updatedTasks)
}

func (s *Service) MarkStatus(id int, status string) error {
	tasks, err := s.repo.Load()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return ErrTaskNotFound
	}

	return s.repo.Save(tasks)
}

func (s *Service) List(filterStatus string) ([]domain.Task, error) {
	tasks, err := s.repo.Load()
	if err != nil {
		return nil, err
	}

	if filterStatus == "" {
		return tasks, nil
	}

	var filtered []domain.Task
	for _, t := range tasks {
		if t.Status == filterStatus {
			filtered = append(filtered, t)
		}
	}
	return filtered, nil
}
