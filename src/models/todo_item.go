package models

import (
	"time"

	"gorm.io/gorm"
)

type TodoItem struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Description string         `json:"description"`
	IsCompleted bool           `json:"is_completed"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}

func NewTodoItem(description string, isCompleted bool) TodoItem {
	return TodoItem{
		Description: description,
		IsCompleted: isCompleted,
	}
}
