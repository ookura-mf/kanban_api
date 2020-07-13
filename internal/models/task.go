package models

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	KanbanID int `json:"kanban_id"`
	Kanban   Kanban
	Title    string `json:"title"`
	Content  string `json:"content"`
	Status   int `json:"status"`
}
