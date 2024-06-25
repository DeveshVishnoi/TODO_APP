package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (string, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loaded the env : ", err)
		return "", err
	}
	return os.Getenv("MONGO_DB_URL"), nil
}
