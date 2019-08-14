package controllers

import (
	"backend/models"
	"database/sql"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func GetTasks(c echo.Context) error {
	var task []models.Task
	var Title string
	var Description string
	var Priority int

	tasksSQL := "select COALESCE(title,'') as title, COALESCE(description,''), COALESCE(priority,0) from task"
	rows, err := Database.Query(tasksSQL)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Title, &Description, &Priority)
		if err != nil {
			log.Println(err)
		}
		a := models.Task{Title:Title, Description:Description, Priority:Priority}
		task = append(task, a)
	}
	return c.JSON(http.StatusOK, task)
}

func GetTask(c echo.Context) error {
	var Title string
	var Description string
	var Priority int

	id := c.Param("id")
	taskSQL := "select COALESCE(title,'') as title, COALESCE(description,''), COALESCE(priority,0) from task where id=?"
	row := Database.QueryRow(taskSQL, id)
	switch err := row.Scan(&Title, &Description, &Priority); err {
	case sql.ErrNoRows:
		log.Println(err)
	case nil:
		log.Println(err)
	}
	task := models.Task{Title:Title,Description:Description,Priority:Priority}
	return c.JSON(http.StatusOK, task)
}

func AddTask(c echo.Context) error {
	var res models.Res
	task := new(models.Task)
	if err := c.Bind(task); err != nil {
		log.Println(err)
	}
	query := "insert into task(title, description, created_date, last_modified, priority) values(?,?,datetime(), datetime(),?)"
	restoreSQL, err := Database.Prepare(query)
	if err != nil {
		log.Println(err)
	}
	tx, err := Database.Begin()
	_, err = tx.Stmt(restoreSQL).Exec(task.Title, task.Description, task.Priority)
	if err != nil {
		log.Print(err)
		res = models.Res{true, "Error"}
	} else {
		log.Printf("insert successful")
		tx.Commit()
		res = models.Res{false,"Task add with success!"}
	}
	return c.JSON(http.StatusCreated, res)
}