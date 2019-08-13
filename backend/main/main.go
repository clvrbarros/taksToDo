package main

import (
	"backend/controllers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/addtask", controllers.AddTask)
	http.HandleFunc("/tasks", controllers.GetTasks)
	fmt.Println("running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
