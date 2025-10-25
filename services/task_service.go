package services

import (
	"errors"
	"strings"

	"github.com/tenuser/myapp/models"
	"github.com/tenuser/myapp/repositories"
)

// TaskService provides business logic for working with tasks.
type TaskService interface {
	ListTasks() ([]models.Task, error)
	GetTask(id int) (models.Task, error)
	CreateTask(input models.Task) (models.Task, error)
	UpdateTask(id int, input models.Task) (models.Task, error)
	DeleteTask(id int) error
}

type taskService struct {
	repo repositories.TaskRepository
}

// NewTaskService wires a repository into a service instance.
func NewTaskService(repo repositories.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) ListTasks() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *taskService) GetTask(id int) (models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) CreateTask(input models.Task) (models.Task, error) {
	if strings.TrimSpace(input.Title) == "" {
		return models.Task{}, errors.New("title is required")
	}

	input.Completed = false
	return s.repo.Create(input)
}

func (s *taskService) UpdateTask(id int, input models.Task) (models.Task, error) {
	if strings.TrimSpace(input.Title) == "" {
		return models.Task{}, errors.New("title is required")
	}

	return s.repo.Update(id, input)
}

func (s *taskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}
