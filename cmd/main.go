package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func getOSEnv(variable string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error encountered loading env file")
	}
	return os.Getenv(variable)
}

func main() {
	db_user := getOSEnv("DB_USER")
	db_password := getOSEnv("DB_PASSWORD")
	db_host := getOSEnv("DB_HOST")
	db_port := getOSEnv("DB_PORT")
	db_name := getOSEnv("DB_NAME")
}