package db

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"
	"os"
	"github.com/vdonkor/vcd-reports/pkg/api/config"
)
var (
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")
)

func DbConnect() *gorm.DB {

	config.LoadOSEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",dbUser,dbPassword,dbHost,dbPort,dbName)
	db, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Printf("Connection established to the %s db\n", dbName)
	return db
}
