package handlers

import (
	"context"
	"studyCRUD/internal/taskService"
	"studyCRUD/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewTaskHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var response []tasks.Task

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   tsk.Task,
			IsDone: tsk.IsDone,
		}
		response = append(response, task)
	}

	return tasks.GetTasks200JSONResponse(response), nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Task:   taskRequest.Task,
		IsDone: taskRequest.IsDone,
		UserID: taskRequest.UserId,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   createdTask.Task,
		IsDone: createdTask.IsDone,
		UserId: createdTask.UserID,
	}

	return response, nil
}

func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := uint(request.Id)

	existingTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var taskToUpdate *taskService.Task
	for i, task := range existingTasks {
		if task.ID == taskID {
			taskToUpdate = &existingTasks[i]
			break
		}
	}

	if taskToUpdate == nil {
		return tasks.PatchTasksId404Response{}, nil
	}

	if request.Body.Task != nil {
		taskToUpdate.Task = request.Body.Task
	}
	if request.Body.IsDone != nil {
		taskToUpdate.IsDone = request.Body.IsDone
	}
	if request.Body.UserId != nil {
		taskToUpdate.UserID = request.Body.UserId
	}

	updatedTask, err := h.Service.UpdateTaskByID(taskID, *taskToUpdate)
	if err != nil {
		return nil, err
	}
	return tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   updatedTask.Task,
		IsDone: updatedTask.IsDone,
		UserId: updatedTask.UserID,
	}, nil
}

func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := uint(request.Id)

	err := h.Service.DeleteTaskByID(taskID)
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}
