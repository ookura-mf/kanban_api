package models

import "github.com/jinzhu/gorm"

type Kanban struct {
	gorm.Model
	Title string
}
