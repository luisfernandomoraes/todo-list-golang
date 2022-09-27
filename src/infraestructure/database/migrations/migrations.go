package migrations

import (
	"github.com/luisfernandomoraes/todo-list-golang/infraestructure/logger"
	"github.com/luisfernandomoraes/todo-list-golang/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	autoMigrateModel(&models.TodoItem{}, db)
}

func autoMigrateModel(model interface{}, db *gorm.DB) {
	if err := db.AutoMigrate(model); err != nil {
		logger.GetLogger().Fatal().Err(err).Str("use-case", "infra/database").Msg("Error in DB migration")
	} else {
		logger.GetLogger().Info().Err(err).Str("use-case", "infra/database").Msg("Success DB migration")
	}
}
