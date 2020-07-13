package main

import (
	"net/http"
	"github.com/kanban_api/internal/controllers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	db, err := gorm.Open("mysql", "root:@/kanban_development?parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Routing
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"health": "ok"})
	})
	e.GET("/kanbans", controllers.GetAllKanbans(db))
	e.GET("/kanbans/:id", controllers.GetKanban(db))
	e.POST("/kanbans", controllers.CreateKanban(db))
	e.PUT("/kanbans/:id", controllers.UpdateKanban(db))
	e.DELETE("/kanbans/:id", controllers.DeleteKanban(db))

	e.GET("/tasks", controllers.GetAllTasks(db))
	e.GET("/tasks/:id", controllers.GetTask(db))
	e.POST("/tasks", controllers.CreateTask(db))
	e.PUT("/tasks/:id", controllers.UpdateTask(db))
	e.DELETE("/tasks/:id", controllers.DeleteTask(db))

	e.Logger.Debug(e.Start(":1313"))
}
