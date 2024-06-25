package main

import (
	"fmt"
	"log"
	db "todo/Db"
	"todo/routers"
	"todo/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	mongoURL, err := utils.LoadEnv()
	if err != nil {
		fmt.Println("Error loading the env file")
		return
	}

	mongoManager, err := db.ConnectDB(mongoURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := gin.Default()
	routers.InitializeRoutes(router, mongoManager)

	router.Run("localhost:5000")
}
