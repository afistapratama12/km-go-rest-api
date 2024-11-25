package main

import (
	"testrestapi/handler"
	"testrestapi/middleware"
	"testrestapi/repository"
	"testrestapi/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// router, inisialisasi gin

	// clean architecture
	// ngembikin 10 data dummy kosong

	taskRepo := repository.NewTaskRepo()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	r := gin.Default()

	// handler (+), service (+), repository (.)

	v1Public := r.Group("/v1")
	{
		v1Public.POST("/login", taskHandler.Login)
	}

	v1 := r.Group("/v1/task", middleware.Auth())
	{
		v1.GET("/", taskHandler.GetTasks)
		v1.GET("/:task-id", taskHandler.GetTask)
		v1.POST("/", taskHandler.CreateTask)
		v1.PUT("/:task-id", taskHandler.UpdateTask)
		v1.DELETE("/:task-id", taskHandler.DeleteTask)
	}

	r.Run(":8080")
}

// auth middleware
// jwt

// task tracker
// api, get list, bikin task list, update, delete

// postman -> dokumentasi api/ call api (clear)
// gin -> framework api (clear)

// clean architecture -> codingan layer
// layer: handler / api / controller [implement]
// layer: service / usecase [implement]
// layer: repository
// layer: model / entity [implement]

// 1 sifat yang sama
