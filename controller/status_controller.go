package controller

import (
	"net/http"
	"project/config"
	"project/model/response"
	"project/model/status"

	"github.com/labstack/echo/v4"
)

func AddStatusReservationCode(c echo.Context) error {
	var status status.Status
	c.Bind(&status)

	make := config.DB.Create(&status)

	if make.Error != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Bad Request - Code sudah digunakan",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "OK- Code berhasil dibuat",
		Data:    nil,
	})

}
