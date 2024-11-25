package service

import (
	"errors"
	"testrestapi/middleware"
	"testrestapi/model"
	"testrestapi/repository"
)

// data structure,

// 2 layer, business logic

// interface, berisi kontrak yang sudah dibikin
type TaskService interface {
	GetTasks() []model.Task
	GetTask(taskID int) model.Task
	CreateTask(request model.TaskRequest) model.Task
	UpdateTask(taskID int, request model.TaskRequest) (model.Task, error)
	DeleteTask(taskID int) error

	LoginProcess(request model.LoginRequest) (string, error)
}

// keep private
type taskService struct {
	TaskRepo repository.TaskRepo
}

func NewTaskService(taskRepo repository.TaskRepo) TaskService {
	return &taskService{
		TaskRepo: taskRepo,
	}
}

func (s *taskService) LoginProcess(request model.LoginRequest) (string, error) {
	// logic login
	// validasi username dan password
	// username= admin, password: admin123

	if request.Username != "admin" || request.Password != "admin123" {
		return "", errors.New("invalid username or password")
	}

	// generate token
	token, err := middleware.GenerateToken(request.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// 1 service untuk forward data ke repo, mendapatkan data dari index

func (s *taskService) GetTasks() []model.Task {
	return s.TaskRepo.GetTasks()
}

func (s *taskService) GetTask(taskID int) model.Task {
	tasks := s.TaskRepo.GetTasks()
	return tasks[taskID]
}

func (s *taskService) CreateTask(request model.TaskRequest) model.Task {
	// login mengolah datanya dan disimpan ke dalam task
	return s.TaskRepo.CreateTask(request)
}

func (s *taskService) UpdateTask(taskID int, request model.TaskRequest) (model.Task, error) {
	tasks := s.TaskRepo.GetTasks()

	// find task id nya di list of tasks
	for i := 0; i < len(tasks); i++ {

		// update datanya sesuai request
		if tasks[i].ID == taskID {
			// update data

			task := s.TaskRepo.UpdateTask(i, request)
			return task, nil
		}
	}

	return model.Task{}, errors.New("task not found")
}

func (s *taskService) DeleteTask(taskID int) error {
	tasks := s.TaskRepo.GetTasks()

	for i := 0; i < len(tasks); i++ {

		// update datanya sesuai request
		if tasks[i].ID == taskID {
			// hapus data
			s.TaskRepo.DeleteTask(i)
			return nil
		}
	}

	return errors.New("task not found")
}
