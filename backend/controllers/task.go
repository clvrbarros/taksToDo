package controllers

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func AddTask (w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var task models.Task
		var res models.Res

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal(body, &task)
		if err != nil {
			fmt.Println(err)
		}

		query := "insert into task(title, description, created_date, last_modified, priority) values(?,?,datetime(), datetime(),?)"
		restoreSQL, err := Database.Prepare(query)
		if err != nil {
			fmt.Println(err)
		}

		tx, err := Database.Begin()
		_, err = tx.Stmt(restoreSQL).Exec(task.Title, task.Description, task.Priority)
		if err != nil {
			fmt.Println(err)
			res = models.Res{true, "Error"}
		} else {
			log.Printf("insert successful")
			tx.Commit()
			res = models.Res{false,"Task add with success!"}
		}

		resJSON, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resJSON)
	}
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var task []models.Task
		var Title string
		var Description string
		var Priority int

		tasksSQL := "select COALESCE(title,'') as title, COALESCE(description,''), COALESCE(priority,0) from task"
		rows, err := Database.Query(tasksSQL)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&Title, &Description, &Priority)
			if err != nil {
				fmt.Println(err)
			}
			a := models.Task{Title:Title, Description:Description, Priority:Priority}
			task = append(task, a)
		}
		contextJSON, err := json.Marshal(task)
		if err != nil {
			log.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(contextJSON)
	}
}