package handler

import (
	"strconv"
	"testrestapi/model"
	"testrestapi/service"

	"github.com/gin-gonic/gin"
)

// js, node js, API express JS
// php, laravel
// java, spring boot

// MVC , model (data), view (tampilan, request response), controller (logic)

// repository, handler, service

// encapsulate
type taskHandler struct {
	// requirement field, layer lain
	TaskService service.TaskService
}

// fungsi init (call public)
func NewTaskHandler(taskService service.TaskService) *taskHandler {
	return &taskHandler{
		TaskService: taskService,
	}
}

func (h *taskHandler) Login(c *gin.Context) {
	var request model.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := h.TaskService.LoginProcess(request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// service melakukan pengecekan data, validasi user pass, generate token

	c.JSON(200, gin.H{"message": "login success", "token": token})
}

func (h *taskHandler) GetTasks(c *gin.Context) {
	tasks := h.TaskService.GetTasks()
	c.JSON(200, tasks)
}

// bikin 1 handler lagi
func (h *taskHandler) GetTask(c *gin.Context) {
	taskIDReq := c.Param("task-id")

	taskID, err := strconv.Atoi(taskIDReq)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	task := h.TaskService.GetTask(taskID)

	c.JSON(200, task)
}

// codingan disini, mengolah request sama response

func (h *taskHandler) CreateTask(c *gin.Context) {
	var request model.TaskRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	task := h.TaskService.CreateTask(request)

	c.JSON(201, gin.H{"message": "task created", "task": task})
}

func (h *taskHandler) UpdateTask(c *gin.Context) {
	taskIDReq := c.Param("task-id")

	taskID, err := strconv.Atoi(taskIDReq)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var request model.TaskRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	task, err := h.TaskService.UpdateTask(taskID, request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "task updated", "task": task})
}

func (h *taskHandler) DeleteTask(c *gin.Context) {
	taskIDReq := c.Param("task-id")

	taskID, err := strconv.Atoi(taskIDReq)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.TaskService.DeleteTask(taskID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "task deleted"})
}
