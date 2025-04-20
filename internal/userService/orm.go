package userService

import (
	"gorm.io/gorm"
	"studyCRUD/internal/web/tasks"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Tasks    []tasks.Task
}
