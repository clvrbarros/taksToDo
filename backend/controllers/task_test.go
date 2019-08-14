package controllers

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddTask(t *testing.T) {
	content := `{
	 "title":"teste3",
     "description":"teste3",
     "priority":3
	}`
	expected := `{"err":false,"message":"Task add with success!"}
`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/task",strings.NewReader(content))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req,rec)
	if assert.NoError(t, AddTask(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestGetTasks(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet,"/tasks",nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req,rec)
	if assert.NoError(t, GetTasks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
