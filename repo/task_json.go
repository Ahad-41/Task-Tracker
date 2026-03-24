package repo

import (
	"encoding/json"
	"os"
	"task-cli/domain"
)

type TaskJSONRepo struct {
	filename string
}

func NewTaskJSONRepo(filename string) *TaskJSONRepo {
	return &TaskJSONRepo{filename: filename}
}

func (r *TaskJSONRepo) Load() ([]domain.Task, error) {
	if _, err := os.Stat(r.filename); os.IsNotExist(err) {
		return []domain.Task{}, nil
	}

	data, err := os.ReadFile(r.filename)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []domain.Task{}, nil
	}

	var tasks []domain.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskJSONRepo) Save(tasks []domain.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filename, data, 0644)
}
