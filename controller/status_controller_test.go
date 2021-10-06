package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"project/config"
	"project/model/status"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InsertDataStatus() error {

	status := status.Status{
		Id:      1,
		Message: "Menunggu Konfirmasi Admin",
	}

	var err error
	if err = config.DB.Save(&status).Error; err != nil {
		return err
	}
	return nil
}
func TestAddStatusUserReservationsCodeSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	requestBody := strings.NewReader(`{"id":1,"message":"Menunggu konfirmasi admin"}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/admin/status", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	if assert.NoError(t, AddStatusReservationCode(c)) {
		response := recorder.Result()
		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "OK", responseBody["message"])
	}
}

func TestAddStatusUserReservationsCodeFailed(t *testing.T) {
	e := InitEchoTestAPI()
	InsertDataStatus()
	requestBody := strings.NewReader(`{"id":1,"message":"Menunggu konfirmasi admin"}`)
	request := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8000/admin/status", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	if assert.NoError(t, AddStatusReservationCode(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Bad Request - Code sudah digunakan", responseBody["message"])
	}
}
