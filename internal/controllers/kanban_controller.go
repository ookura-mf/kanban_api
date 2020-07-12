package controllers

import (
	"net/http"
	"github.com/kanban_api/internal/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetAllKanbans(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var kanbans []models.Kanban
		db.Find(&kanbans)
		return c.JSON(http.StatusOK, kanbans)
	}
}

func CreateKanban(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		kanban := new(models.Kanban)
		if err := c.Bind(kanban); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"messages": "something wrong!"})
		}
		db.Create(&kanban)
	
		return c.JSON(http.StatusOK, kanban)
	}
}
