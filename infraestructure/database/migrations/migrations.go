package migrations

import (
	"github.com/luisfernandomoraes/todo-list-golang/models"
	"gorm.io/gorm"
	"log"
)

func RunMigrations(db *gorm.DB) {
	autoMigrateModel(&models.TodoItem{}, db)
}

func autoMigrateModel(model interface{}, db *gorm.DB) {
	if err := db.AutoMigrate(model); err != nil {
		log.Fatalf("Error in DB migration: %s", err.Error())
	}
}
