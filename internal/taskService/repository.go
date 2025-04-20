package taskService

import (
	"errors"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)

	GetAllTasks() ([]Task, error)

	GetTaskByID(id uint) (Task, error)

	GetTasksByUserID(userID uint) ([]Task, error)

	UpdateTaskByID(id uint, task Task) (Task, error)

	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (repo *taskRepository) CreateTask(task Task) (Task, error) {
	result := repo.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (repo *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := repo.db.Find(&tasks).Error
	return tasks, err
}

func (repo *taskRepository) GetTaskByID(id uint) (Task, error) {
	var tasks Task
	result := repo.db.First(&tasks, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Task{}, ErrTaskNotFound
		}
		return Task{}, result.Error
	}
	return tasks, nil
}

func (repo *taskRepository) GetTasksByUserID(userID uint) ([]Task, error) {
	var dbTasks []Task
	err := repo.db.Where("user_id = ?", userID).Find(&dbTasks).Error
	if err != nil {
		return nil, err
	}

	return dbTasks, nil
}

func (repo *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	var existingTask Task

	if err := repo.db.First(&existingTask, id).Error; err != nil {
		return Task{}, err
	}

	result := repo.db.Model(&existingTask).Updates(task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return existingTask, nil
}

func (repo *taskRepository) DeleteTaskByID(id uint) error {
	result := repo.db.Where("id = ?", id).Delete(&Task{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
