package userService

import (
	"errors"
	"regexp"
	"studyCRUD/internal/web/tasks"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) GetAllUsers() ([]User, error) {
	return service.repo.GetAllUsers()
}

func (service *UserService) GetTasksForUser(userID uint) ([]tasks.Task, error) {
	return service.repo.GetTasksForUser(userID)
}

func isValidUsername(name string) bool {
	if len(name) < 3 || len(name) > 20 {
		return false
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return re.MatchString(name)
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	re := regexp.MustCompile(`^(?=.*[a-zA-Z])(?=.*[0-9])(?=.*[^a-zA-Z0-9]).{8,}$`)
	return re.MatchString(password)
}

func (service *UserService) CreateUser(user User) (User, error) {

	if !isValidUsername(user.Name) {
		return User{}, errors.New("invalid username: must be 3-20 characters, letters/numbers only")
	}
	if !isValidEmail(user.Email) {
		return User{}, errors.New("invalid email format")
	}
	if !isValidPassword(user.Password) {
		return User{}, errors.New("invalid password: must be at least 8 characters with letters, numbers, and special symbols")
	}

	existingUser, err := service.repo.GetUserByEmail(user.Email)
	if err == nil && existingUser.ID != 0 {
		return User{}, errors.New("user with this email already exists")
	}

	return service.repo.CreateUser(user)
}

func (service *UserService) UpdateUserByID(id uint, user User) (User, error) {
	return service.repo.UpdateUserByID(id, user)
}

func (service *UserService) DeleteUserByID(id uint) error {
	return service.repo.DeleteUserByID(id)
}
