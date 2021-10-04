package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {

}
func InsertBookData(e *echo.Echo) error {
	requestBody := strings.NewReader(`{"id":"XfFvDwAAQBAJ","qty":"10"}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/addBooks", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	return AddBookController(c)
}
func TestAddBookControllerSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	// InsertBookData(e)
	requestBody := strings.NewReader(`{"id":"XfFvDwAAQBAJ","qty":10}`)
	request := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8000/user/login", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, AddBookController(c)) {
		response := recorder.Result()
		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "Buku berhasil ditambahkan", responseBody["message"])
		assert.Equal(t, "Atomic Habits", responseBody["data"].(map[string]interface{})["title"])
	}

}
