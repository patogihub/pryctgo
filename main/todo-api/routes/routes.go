package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/patogihub/pryctgo/todo-api/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	t := r.Group("/tasks")
	{
		t.GET("/", handler.GetTasks)
		t.POST("/", handler.CreateTask)
		t.PUT("/:id", handler.UpdateTask)
		t.DELETE("/:id", handler.DeleteTask)
	}

	return r
}
