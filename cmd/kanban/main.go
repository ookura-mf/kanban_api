package main

import (
	"net/http"

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

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"health": "ok"})
	})

	e.Logger.Fatal(e.Start(":1313"))
}
