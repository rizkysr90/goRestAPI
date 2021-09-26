package controller

import (
	"errors"
	"net/http"
	"project/config"
	"project/middlewares"
	admins "project/model/admin"
	"project/model/response"

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
