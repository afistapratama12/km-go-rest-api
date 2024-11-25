package repository

import "testrestapi/model"

// layer: data, update data, delete data, get data

var Tasks = []model.Task{}
var TotalTask = 0

type TaskRepo interface {
	GetTasks() []model.Task
	CreateTask(request model.TaskRequest) model.Task
	UpdateTask(idx int, request model.TaskRequest) model.Task
	DeleteTask(idx int)
}

// keep private
type taskRepo struct {
}

func NewTaskRepo() TaskRepo {
	return &taskRepo{}
}

func (r *taskRepo) GetTasks() []model.Task {
	return Tasks
}

func (r *taskRepo) CreateTask(request model.TaskRequest) model.Task {
	TotalTask++

	task := model.Task{
		ID:    TotalTask,
		Title: request.Title,
		Body:  request.Body,
	}

	Tasks = append(Tasks, task)
	return task
}

func (r *taskRepo) UpdateTask(idx int, request model.TaskRequest) model.Task {
	Tasks[idx].Title = request.Title
	Tasks[idx].Body = request.Body
	return Tasks[idx]
}

func (r *taskRepo) DeleteTask(idx int) {
	Tasks = append(Tasks[:idx], Tasks[idx+1:]...)
}
