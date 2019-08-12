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
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		var task models.Task
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
		} else {
			log.Printf("insert successful")
			tx.Commit()
		}

		res := models.Res{false,"Task add with success!"}
		resJSON, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
		}

		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resJSON)
	}
}