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
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/addBooks", requestBody)
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

func TestAddBookControllerFailedDuplicate(t *testing.T) {
	e := InitEchoTestAPI()
	InsertBookData(e)
	requestBody := strings.NewReader(`{"id":"XfFvDwAAQBAJ","qty":10}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/addBooks", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, AddBookController(c)) {
		response := recorder.Result()
		assert.Equal(t, 500, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 500, int(responseBody["code"].(float64)))
		assert.Equal(t, "Buku sudah tersedia sebelumnya", responseBody["message"])
	}
}

func TestAddBookControllerFailed(t *testing.T) {
	e := InitEchoTestAPI()
	// InsertBookData(e)
	requestBody := strings.NewReader(`{"id":"","qty":10}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/addBooks", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, AddBookController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Required parameter", responseBody["message"])
	}
}
func TestAddBookControllerFailByVolumeId(t *testing.T) {
	e := InitEchoTestAPI()
	// InsertBookData(e)
	requestBody := strings.NewReader(`{"id":"asdhjadkahdkahdkadha","qty":10}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/addBooks", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, AddBookController(c)) {
		response := recorder.Result()
		assert.Equal(t, 204, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 204, int(responseBody["code"].(float64)))
		assert.Equal(t, "Oops,Something wrong", responseBody["message"])
	}
}

func TestSearchBookByTitleSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	InsertBookData(e)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/search?title=Atomic", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, SearchBookByTitle(c)) {
		response := recorder.Result()
		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "OK", responseBody["message"])
	}

}
func TestSearchBookByTitleFailed(t *testing.T) {
	e := InitEchoTestAPI()
	InsertBookData(e)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/search?title=Universe", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, SearchBookByTitle(c)) {
		response := recorder.Result()
		assert.Equal(t, 204, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 204, int(responseBody["code"].(float64)))
		assert.Equal(t, "Data Buku Tidak Ditemukan", responseBody["message"])
	}
}
