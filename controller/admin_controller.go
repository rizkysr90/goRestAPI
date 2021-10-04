package controller

import (
	"errors"
	"net/http"
	"project/config"
	"project/helper"
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
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request - Nama wajib diisi saat pendaftaran",
			Data:    nil,
		})
	}
	if adminRegister.Email == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request - Email wajib diisi saat pendaftaran",
			Data:    nil,
		})
	}
	if adminRegister.Password == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request - Password wajib diisi saat pendaftaran",
			Data:    nil,
		})
	}

	var AdminDB admins.Admin
	AdminDB.Name = adminRegister.Name
	AdminDB.Email = adminRegister.Email
	AdminDB.Address = adminRegister.Address
	AdminDB.Password, _ = helper.Hash(adminRegister.Password)

	result := config.DB.Create(&AdminDB)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request - Email sudah digunakan",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    nil,
	})
}

func LoginAdminController(c echo.Context) error {
	adminLogin := admins.AdminLogin{}
	c.Bind(&adminLogin)

	admin := admins.Admin{}

	result := config.DB.First(&admin, "email = ?", adminLogin.Email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusForbidden, response.BaseResponse{
				Code:    http.StatusForbidden,
				Message: "Periksa email anda kembali",
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
	if !helper.CheckPasswordHash(adminLogin.Password, admin.Password) {
		return c.JSON(http.StatusForbidden, response.BaseResponse{
			Code:    http.StatusForbidden,
			Message: "Password salah",
			Data:    nil,
		})
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
