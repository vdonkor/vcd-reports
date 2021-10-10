package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)
func LoadOSEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error encountered loading env file")
	}
}