package main

import (
	"backend/controllers"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func main() {
	defer controllers.Database.Close()
	e := echo.New()
	e.GET("/task/:id", controllers.GetTask)
	e.GET("/tasks", controllers.GetTasks)
	e.POST("/task", controllers.AddTask)
	log.Fatal(http.ListenAndServe(":8080",e))
}
