package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/luisfernandomoraes/todo-list-golang/infraestructure/database/migrations"
	logger "github.com/luisfernandomoraes/todo-list-golang/infraestructure/logs"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		// logger.GetLogger().WithFields(logrus.Fields{
		// 	"environment": serviceconfig.Environment,
		// 	"use-case":    "load-database",
		// 	"type":        "technical",
		// }).Error(err.Error())
		uuid := uuid.New()
		logger.Fatal("An error ocurred at open conection to PostgresDB.", zap.Error(err), zap.String("guid", uuid.String()))
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)

	// logger.GetLogger().WithFields(logrus.Fields{
	// 	"environment": serviceconfig.Environment,
	// 	"use-case":    "load-database",
	// 	"type":        "technical",
	// }).Info("Loaded database successfully")
	log.Println("Loaded database successfully")
}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}

func SetDatabase(DB *gorm.DB) {
	db = DB
}
