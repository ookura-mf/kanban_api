package controllers

import (
	"net/http"
	"github.com/kanban_api/internal/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetAllTasks(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tasks []models.Task
		db.Find(&tasks)
		return c.JSON(http.StatusOK, tasks)
	}
}

func GetTask(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		if id := c.Param("id"); id != "" {
			var task models.Task
			db.Find(&task, id)
			if task.ID != 0 {
				return c.JSON(http.StatusOK, task)
			}
		}

		return c.JSON(http.StatusNotFound, nil)
	}
}

func CreateTask(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		task := new(models.Task)
		if err := c.Bind(task); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"messages": "something wrong!"})
		}
		task.Status = 1
		db.Create(&task)
	
		return c.JSON(http.StatusOK, task)
	}
}

func UpdateTask(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		newTask := new(models.Task)
		if err := c.Bind(newTask); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"messages": "something wrong!"})
		}

		if id := c.Param("id"); id != "" {
			var task models.Task
			db.Find(&task, id)
			if task.ID != 0 {
				if newTask.Title != "" {
					task.Title = newTask.Title
				}
				if newTask.Content != "" {
					task.Content = newTask.Content
				}
				if newTask.Status != 0 {
					task.Status = newTask.Status
				}
				db.Save(&task)
				return c.JSON(http.StatusOK, task)
			}
		}
		return c.JSON(http.StatusNotFound, nil)
	}
}

func DeleteTask(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		if id := c.Param("id"); id != "" {
			var task models.Task
			db.Find(&task, id)
			if task.ID != 0 {
				db.Delete(task)
				return c.JSON(http.StatusOK, task)
			}
		}
		return c.JSON(http.StatusNotFound, nil)
	}
}
