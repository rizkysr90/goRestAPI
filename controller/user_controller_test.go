package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"project/config"
	"project/helper"
	"project/model/users"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func InsertDataUser() error {

	passwordori, _ := helper.Hash("adarizki123")
	user := users.User{
		Name:     "Rizki",
		Password: passwordori,
		Email:    "rizkitest123@gmail.com",
	}

	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
func TestRegisterUserControllerSuccess(t *testing.T) {
	e := InitEchoTestAPI() //make a router with echo

	requestBody := strings.NewReader(`{"name":"Rizky","email":"rizkytesting@gmail.com","password":"lagitest123","address":"bekasi"}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/user/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterUserController(c)) {
		response := recorder.Result()
		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "OK", responseBody["message"])
		// assert.Equal(t, http.StatusOK, recorder.Code)
		// response := rec.Body.String()
	}
	// response := recorder.Result()

	// assert.Equal(t, 200, response.StatusCode)
	// assert.Equal(t, "OK", response.Status)
	// assert.Equal(t, nil, responseBody["data"].(map[string]interface{})["name"])
}

func TestRegisterUserControllerFailedName(t *testing.T) {
	e := InitEchoTestAPI() //make a router with echo
	requestBody := strings.NewReader(`{"name":"","email":"rizkytesting@gmail.com","password":"lagitest123","address":"bekasi"}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/user/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterUserController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Bad Request - Nama wajib diisi saat pendaftaran", responseBody["message"])
		// assert.Equal(t, http.StatusOK, recorder.Code)
		// response := rec.Body.String()
	}
}

func TestRegisterUserControllerFailedEmail(t *testing.T) {
	e := InitEchoTestAPI() //make a router with echo
	requestBody := strings.NewReader(`{"name":"rizky","email":"","password":"lagitest123","address":"bekasi"}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/user/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterUserController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Bad Request - Email wajib diisi saat pendaftaran", responseBody["message"])
		// assert.Equal(t, http.StatusOK, recorder.Code)
		// response := rec.Body.String()
	}
}
func TestRegisterUserControllerFailedPassword(t *testing.T) {
	e := InitEchoTestAPI() //make a router with echo
	requestBody := strings.NewReader(`{"name":"rizky","email":"emailnyaada@gmail.com","password":"","address":"bekasi"}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/user/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterUserController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Bad Request - Password wajib diisi saat pendaftaran", responseBody["message"])
		// assert.Equal(t, http.StatusOK, recorder.Code)
		// response := rec.Body.String()
	}
}
func TestRegisterUserControllerFailedDuplicatEmail(t *testing.T) {
	e := InitEchoTestAPI() //make a router with echo
	InsertDataUser()
	requestBody := strings.NewReader(`{"name":"Rizki","email":"rizkitest123@gmail.com","password":"adarizki123","address":"bekasi"}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/user/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterUserController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Bad Request - Email sudah digunakan", responseBody["message"])
		// assert.Equal(t, http.StatusOK, recorder.Code)
		// response := rec.Body.String()
	}
}

func TestLoginUserControllerSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	InsertDataUser()

	requestBody := strings.NewReader(`{"name":"Rizki","email":"rizkitest123@gmail.com","password":"adarizki123","address":"bekasi"}`)
	request := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8000/user/login", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, LoginUserController(c)) {
		response := recorder.Result()
		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "Hi,Welcome back", responseBody["message"])
		assert.Equal(t, "Rizki", responseBody["data"].(map[string]interface{})["name"])
		assert.Equal(t, "rizkitest123@gmail.com", responseBody["data"].(map[string]interface{})["email"])
	}
}

func TestLoginUserControllerFailedByEmail(t *testing.T) {
	e := InitEchoTestAPI()
	InsertDataUser()

	requestBody := strings.NewReader(`{"name":"Rizki","email":"rizkitest1234@gmail.com","password":"adarizki123","address":"bekasi"}`)
	request := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8000/user/login", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, LoginUserController(c)) {
		response := recorder.Result()
		assert.Equal(t, 403, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 403, int(responseBody["code"].(float64)))
		assert.Equal(t, "Email belum terdaftar", responseBody["message"])
	}
}
func TestLoginUserControllerFailedByPassword(t *testing.T) {
	e := InitEchoTestAPI()
	InsertDataUser()

	requestBody := strings.NewReader(`{"name":"Rizki","email":"rizkitest1234@gmail.com","password":"adarizki123","address":"bekasi"}`)
	request := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8000/user/login", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, LoginUserController(c)) {
		response := recorder.Result()
		assert.Equal(t, 403, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 403, int(responseBody["code"].(float64)))
		assert.Equal(t, "Email belum terdaftar", responseBody["message"])
	}
}
