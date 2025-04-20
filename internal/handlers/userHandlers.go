package handlers

import (
	"context"
	"errors"
	"studyCRUD/internal/userService"
	"studyCRUD/internal/web/users"
)

type Handlers struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *Handlers { return &Handlers{Service: service} }

func (h *Handlers) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}

	for _, user := range allUsers {
		user := users.User{
			Id:    &user.ID,
			Email: &user.Email,
			Name:  &user.Name,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *Handlers) GetUsersIdTasks(_ context.Context, request users.GetUsersIdTasksRequestObject) (users.GetUsersIdTasksResponseObject, error) {
	userID := request.Id

	dbTasks, err := h.Service.GetTasksForUser(uint(userID))
	if err != nil {
		return nil, err
	}

	var responseTasks []users.Task
	for _, task := range dbTasks {
		responseTask := users.Task{
			Id:     task.Id,
			Task:   "",
			IsDone: task.IsDone,
			UserId: 0,
		}
		if task.UserId != nil {
			responseTask.UserId = *task.UserId
		}
		if task.Task != nil {
			responseTask.Task = *task.Task
		}

		responseTasks = append(responseTasks, responseTask)
	}

	return users.GetUsersIdTasks200JSONResponse(responseTasks), nil
}

func (h *Handlers) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := userService.User{

		Email:    *userRequest.Email,
		Name:     *userRequest.Name,
		Password: *userRequest.Password,
	}
	createdUser, err := h.Service.CreateUser(userToCreate)

	if userRequest.Password == nil {
		return nil, errors.New("password is required")
	}

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:    &createdUser.ID,
		Name:  &createdUser.Name,
		Email: &createdUser.Email,
	}

	return response, nil
}

func (h *Handlers) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userID := uint(request.Id)

	existingUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var userToUpdate *userService.User
	for i, user := range existingUsers {
		if user.ID == userID {
			userToUpdate = &existingUsers[i]
			break
		}
	}

	if userToUpdate == nil {
		return users.PatchUsersId404Response{}, nil
	}
	if request.Body.Name != nil {
		userToUpdate.Name = *request.Body.Name
	}
	if request.Body.Email != nil {
		userToUpdate.Email = *request.Body.Email
	}
	if request.Body.Password != nil {
		userToUpdate.Password = *request.Body.Password
	}

	updatedUser, err := h.Service.UpdateUserByID(userID, *userToUpdate)
	if err != nil {
		return nil, err
	}
	return users.PatchUsersId200JSONResponse{
		Id:       &updatedUser.ID,
		Name:     &updatedUser.Name,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}, nil
}

func (h *Handlers) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userID := uint(request.Id)

	err := h.Service.DeleteUserByID(userID)
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}
