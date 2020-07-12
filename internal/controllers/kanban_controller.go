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

func GetKanban(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		if id := c.Param("id"); id != "" {
			var kanban models.Kanban
			db.Find(&kanban, id)
			if kanban.ID != 0 {
				return c.JSON(http.StatusOK, kanban)
			}
		}

		return c.JSON(http.StatusNotFound, nil)
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

func UpdateKanban(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		newKanban := new(models.Kanban)
		if err := c.Bind(newKanban); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"messages": "something wrong!"})
		}

		if id := c.Param("id"); id != "" {
			var kanban models.Kanban
			db.Find(&kanban, id)
			if kanban.ID != 0 {
				kanban.Title = newKanban.Title
				db.Save(&kanban)
				return c.JSON(http.StatusOK, kanban)	
			}
		}
		return c.JSON(http.StatusNotFound, nil)
	}
}

func DeleteKanban(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		if id := c.Param("id"); id != "" {
			var kanban models.Kanban
			db.Find(&kanban, id)
			if kanban.ID != 0 {
				db.Delete(kanban)
				return c.JSON(http.StatusOK, kanban)
			}
		}
		return c.JSON(http.StatusNotFound, nil)
	}
}
