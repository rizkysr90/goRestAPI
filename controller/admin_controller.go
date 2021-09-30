package controller

import (
	"errors"
	"fmt"
	"net/http"
	"project/config"
	"project/middlewares"
	admins "project/model/admin"
	"project/model/loan"
	"project/model/response"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterAdminController(c echo.Context) error {
	var adminRegister admins.AdminRegister

	c.Bind(&adminRegister)

	if adminRegister.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Nama tidak boleh kosong!",
		})
	}

	var AdminDB admins.Admin
	AdminDB.Name = adminRegister.Name
	AdminDB.Email = adminRegister.Email
	AdminDB.Address = adminRegister.Address
	AdminDB.Password = adminRegister.Address

	result := config.DB.Create(&AdminDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create the data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes create user",
		"id":      AdminDB.Id,
		"name":    AdminDB.Name,
	})
}

func LoginAdminController(c echo.Context) error {
	adminLogin := admins.AdminLogin{}
	c.Bind(&adminLogin)

	admin := admins.Admin{}

	result := config.DB.First(&admin, "email = ? AND password = ?", adminLogin.Email, adminLogin.Password)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusForbidden, response.BaseResponse{
				Code:    http.StatusForbidden,
				Message: "User tidak ditemukan atau password tidak sesuai",
				Data:    nil,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Ada keselahan di server",
				Data:    nil,
			})
		}
	}
	token, err := middlewares.GenerateTokenJWTAdmin(admin.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Ada keselahan di server",
			Data:    nil,
		})
	}
	adminResponse := admins.AdminResponse{
		Id:        admin.Id,
		Name:      admin.Name,
		Email:     admin.Email,
		Address:   admin.Address,
		Token:     token,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil login",
		Data:    adminResponse,
	})

}

// func GetReservation(c echo.Context) error {

// }
func ReservationProcces(c echo.Context) error {

	getJSON := loan.ProccesReservation{}
	c.Bind(&getJSON)
	reservation := loan.Loan{}
	result := config.DB.First(&reservation, "id = ?", getJSON.Id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusForbidden, response.BaseResponse{
				Code:    http.StatusForbidden,
				Message: "Reservation tidak ditemukan",
				Data:    nil,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Ada keselahan di server",
				Data:    nil,
			})
		}
	}
	layout := "2006-01-02" //TEMPLATE PARSE STRING TO DATE
	reservation.Status = getJSON.Status
	date, err := time.Parse(layout, getJSON.LoanDate)

	if err != nil {
		fmt.Println(err)
	}

	reservation.LoanDate = date
	config.DB.Save(&reservation)

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Pesanan dikonfirmasi",
		Data:    nil,
	})

}
