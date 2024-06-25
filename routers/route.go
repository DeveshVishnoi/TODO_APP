package routers

import (
	db "todo/Db"
	"todo/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	dbHelper    db.DbHelperProvider
	MongoClient db.Manager
}

func InitializeRoutes(router *gin.Engine, dbHelper db.DbHelperProvider) {
	db.DbFunction = dbHelper

	router.POST("/register_user", handlers.CreateUser)
	router.POST("/login", handlers.LoginUser)
	router.GET("/getAllUser", handlers.GetUsers)
	router.POST("/add-task", handlers.CreateTask)
	router.GET("/getAllTask", handlers.GetAllTasks)
	router.PUT("/update-task", handlers.UpdateTask)
	router.GET("/tasks/:user_id", handlers.GetTasksByUser)
	router.DELETE("/tasks/:task_id", handlers.DeleteTask)
}
