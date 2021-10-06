package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"project/config"
	"project/helper"
	admins "project/model/admin"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InsertDataAdmin() error {

	passwordori, _ := helper.Hash("adaipul123")
	admin := admins.Admin{
		Name:     "Syaiful",
		Password: passwordori,
		Email:    "syaifulkece@gmail.com",
	}

	var err error
	if err = config.DB.Save(&admin).Error; err != nil {
		return err
	}
	return nil
}
func TestRegisterAdminControllerSuccess(t *testing.T) {
	e := InitEchoTestAPI() //make a router with echo

	requestBody := strings.NewReader(`{"name":"Syaiful Bahri","email":"syaifultest@gmail.com","password":"lagitest123","address":"bekasi"}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/admin/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterAdminController(c)) {
		response := recorder.Result()
		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "OK", responseBody["message"])
	}
}
func TestRegisterAdminControllerFailedName(t *testing.T) {
	e := InitEchoTestAPI() //make a router with echo
	requestBody := strings.NewReader(`{"name":"","email":"syaifultest@gmail.com","password":"lagitest123","address":"bekasi"}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/admin/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterAdminController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Bad Request - Nama wajib diisi saat pendaftaran", responseBody["message"])
	}
}
func TestRegisterAdminControllerFailedEmail(t *testing.T) {
	e := InitEchoTestAPI() //make a router with echo
	requestBody := strings.NewReader(`{"name":"Syaiful Bahri","email":"","password":"lagitest123","address":"bekasi"}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/admin/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterAdminController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Bad Request - Email wajib diisi saat pendaftaran", responseBody["message"])
	}
}
func TestRegisterAdminControllerFailedPassword(t *testing.T) {
	e := InitEchoTestAPI() //make a router with echo
	requestBody := strings.NewReader(`{"name":"Syaiful","email":"syaifulkece@gmail.com","password":"","address":"bekasi"}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/admin/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterAdminController(c)) {
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
func TestRegisterAdminControllerFailedDuplicatEmail(t *testing.T) {
	e := InitEchoTestAPI() //make a router with echo
	InsertDataAdmin()
	requestBody := strings.NewReader(`{"name":"Syaiful","email":"syaifulkece@gmail.com","password":"adaipul123","address":"bekasi"}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/admin/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterAdminController(c)) {
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
func TestLoginAdminControllerSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	InsertDataAdmin()

	requestBody := strings.NewReader(`{"email":"syaifulkece@gmail.com","password":"adaipul123"}`)
	request := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8000/admin/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, LoginAdminController(c)) {
		response := recorder.Result()
		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "Berhasil login", responseBody["message"])
		assert.Equal(t, "Syaiful", responseBody["data"].(map[string]interface{})["name"])
		assert.Equal(t, "syaifulkece@gmail.com", responseBody["data"].(map[string]interface{})["email"])
	}
}
func TestLoginAdminControllerFailedByEmail(t *testing.T) {
	e := InitEchoTestAPI()
	InsertDataAdmin()

	requestBody := strings.NewReader(`{"email":"syaifulkece123@gmail.com","password":"adaipul123"}`)
	request := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8000/admin/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, LoginAdminController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Periksa email anda kembali", responseBody["message"])
	}
}
func TestLoginAdminControllerFailedByPassword(t *testing.T) {
	e := InitEchoTestAPI()
	InsertDataAdmin()

	requestBody := strings.NewReader(`{"email":"syaifulkece@gmail.com","password":"adaipul123456"}`)
	request := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8000/admin/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, LoginAdminController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Password salah", responseBody["message"])
	}
}
