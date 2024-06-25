package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	db "todo/Db"
	"todo/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func PrintData(c *gin.Context) {
	c.JSON(http.StatusOK, `hello my name is devesh`)
}

func CreateUser(c *gin.Context) {
	var userInput models.User

	err := json.NewDecoder(c.Request.Body).Decode(&userInput)
	if err != nil {
		fmt.Println("Unable to decode the user Input")
		return
	}

	fmt.Println("UserInput", userInput)
	if userInput.EmailId == "" {
		c.JSON(http.StatusBadRequest, "email Id can not be empty")
		return
	}

	if userInput.Password == "" {
		c.JSON(http.StatusBadRequest, "password can not be empty")
		return
	}

	if userInput.Name == "" {
		c.JSON(http.StatusBadRequest, "Name can not be empty")
		return
	}

	userInput.UUID = uuid.New().String()

	userdata, err := db.DbFunction.InsertUser(userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"UserData": userdata})

}

func LoginUser(c *gin.Context) {

	var userInput models.User

	err := json.NewDecoder(c.Request.Body).Decode(&userInput)
	if err != nil {
		fmt.Println("Unable to decode the user Input")
		return
	}

	if userInput.EmailId == "" {
		c.JSON(http.StatusBadRequest, "email Id can not be empty")
		return
	}

	if userInput.Password == "" {
		c.JSON(http.StatusBadRequest, "password can not be empty")
		return
	}

	userData, err := db.DbFunction.GetUser(userInput.EmailId)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, gin.H{
			"userData": "",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	if userData.Password != userInput.Password {

		c.JSON(http.StatusOK, gin.H{
			"UserData": "Password doesn't match...!!!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"UserData": userData,
	})

}

func GetUsers(c *gin.Context) {

	userData, err := db.DbFunction.GetAllUsers()
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, gin.H{
			"userData": "no documents of the user...",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"userData": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userData": userData,
	})

}

func CreateTask(c *gin.Context) {

	var taskInput models.Task

	err := json.NewDecoder(c.Request.Body).Decode(&taskInput)
	if err != nil {
		fmt.Println("Error decode the task", err)
		return
	}

	if taskInput.TaskName == "" {
		c.JSON(http.StatusBadRequest, "task name is not empty...!!!")
		return
	}

	if taskInput.UserId == "" {
		c.JSON(http.StatusBadRequest, "userId is not Empty...!!!")
		return
	}

	taskInput.UUID = uuid.New().String()
	taskInput.TaskDate = time.Now().Unix()

	taskUUID, err := db.DbFunction.InsertTaskInUser(taskInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"taskUUID": "",
			"Error":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"taskUUID": taskUUID,
	})

}

func GetAllTasks(c *gin.Context) {
	taskData, err := db.DbFunction.GetAllTasks()
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, gin.H{
			"taskData": "no documents of the task...",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"taskData": "",
			"Error":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"taskData": taskData,
	})
}

func UpdateTask(c *gin.Context) {

	var taskInput models.Task

	err := json.NewDecoder(c.Request.Body).Decode(&taskInput)
	if err != nil {
		fmt.Println("Error decode the task", err)
		return
	}

	if taskInput.TaskName == "" {
		c.JSON(http.StatusBadRequest, "task name is not empty...!!!")
		return
	}

	if taskInput.UUID == "" {
		c.JSON(http.StatusBadRequest, "UUID is not empty , How can I filter???...!!!")
		return
	}

	if taskInput.UserId == "" {
		c.JSON(http.StatusBadRequest, "userId is not Empty...!!!")
		return
	}

	taskInput.TaskDate = time.Now().Unix()

	taskData, err := db.DbFunction.UpdateTaskInUser(taskInput, taskInput.UUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"taskUUID": "",
			"Error":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"task":   taskData,
		"Status": "Successfully Updated",
	})

}

func GetTasksByUser(c *gin.Context) {

	userId := c.Param("user_id")
	taskData, err := db.DbFunction.GetUserTasks(userId)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, gin.H{
			"taskData": "no documents of the task...",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"taskData": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"taskData": taskData,
	})

}

func DeleteTask(c *gin.Context) {
	task_id := c.Param("task_id")
	data, err := db.DbFunction.DeleteTaskInUser(task_id)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, gin.H{
			"userData": "no documents of the task...",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"taskData": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": data,
	})

}
