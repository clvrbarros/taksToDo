package models

//Task is type of the task
type Task struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Priority int `json:"priority"`
}

//Res is type the model for responses
type Res struct {
	Err string `json:"err"`
	Message string `json:"message"`
}