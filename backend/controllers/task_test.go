package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddTask(t *testing.T) {
	var jsonStr = []byte(`{"title":"teste","description":"tested","priority":0}`)

	req, err := http.NewRequest("POST", "/addtask",bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddTask)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"err":false,"message":"Task add with success!"}`
	if rr.Body.String() != expected {
		t.Errorf("handler func unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetTasks(t *testing.T) {

}