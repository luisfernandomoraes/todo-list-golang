package database

import (
	"fmt"
	"os"
	"time"

	"github.com/luisfernandomoraes/todo-list-golang/infraestructure/database/migrations"
	"github.com/luisfernandomoraes/todo-list-golang/infraestructure/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		logger.GetLogger().Fatal().Err(err).Msg("An error occurred at open connection to PostgresDB.")
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)

	logger.GetLogger().Info().Str("use-case", "infra/database").Msg("Loaded database successfully")
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
