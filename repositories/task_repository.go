package repositories

import (
	"errors"
	"sort"
	"sync"

	"github.com/tenuser/myapp/models"
)

// TaskRepository exposes persistence operations for tasks.
type TaskRepository interface {
	GetAll() ([]models.Task, error)
	GetByID(id int) (models.Task, error)
	Create(task models.Task) (models.Task, error)
	Update(id int, task models.Task) (models.Task, error)
	Delete(id int) error
}

// InMemoryTaskRepository stores tasks in a simple map guarded by a mutex.
type InMemoryTaskRepository struct {
	mu     sync.RWMutex
	tasks  map[int]models.Task
	nextID int
}

// NewInMemoryTaskRepository constructs an empty repository instance.
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

// GetAll returns every task sorted by ID to keep responses deterministic.
func (r *InMemoryTaskRepository) GetAll() ([]models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.tasks) == 0 {
		return []models.Task{}, nil
	}

	result := make([]models.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		result = append(result, task)
	}

	sort.Slice(result, func(i, j int) bool { return result[i].ID < result[j].ID })
	return result, nil
}

// GetByID retrieves a single task by its identifier.
func (r *InMemoryTaskRepository) GetByID(id int) (models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, ok := r.tasks[id]
	if !ok {
		return models.Task{}, errors.New("task not found")
	}

	return task, nil
}

// Create persists a new task and auto-increments the ID.
func (r *InMemoryTaskRepository) Create(task models.Task) (models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task.ID = r.nextID
	r.nextID++
	r.tasks[task.ID] = task

	return task, nil
}

// Update replaces an existing task.
func (r *InMemoryTaskRepository) Update(id int, task models.Task) (models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return models.Task{}, errors.New("task not found")
	}

	task.ID = id
	r.tasks[id] = task
	return task, nil
}

// Delete removes a task by ID.
func (r *InMemoryTaskRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return errors.New("task not found")
	}

	delete(r.tasks, id)
	return nil
}
