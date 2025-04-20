package taskService

import "errors"

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (service *TaskService) CreateTask(task Task) (Task, error) {
	return service.repo.CreateTask(task)
}

func (service *TaskService) GetAllTasks() ([]Task, error) {

	return service.repo.GetAllTasks()
}

func (service *TaskService) GetTaskByID(id uint) (Task, error) {
	return service.repo.GetTaskByID(id)
}

func (service *TaskService) UpdateTaskByID(id uint, task Task) (Task, error) {
	return service.repo.UpdateTaskByID(id, task)
}

func (service *TaskService) DeleteTaskByID(id uint) error {
	return service.repo.DeleteTaskByID(id)
}

var ErrTaskNotFound = errors.New("task not found")
