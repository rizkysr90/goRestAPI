package controller

import (
	"errors"
	"net/http"
	"project/config"
	"project/helper"
	"project/middlewares"
	"project/model/response"
	"project/model/users"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterUserController(c echo.Context) error {
	var userRegister users.UserRegister

	c.Bind(&userRegister)

	if userRegister.Name == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request - Nama wajib diisi saat pendaftaran",
			Data:    nil,
		})
	}
	if userRegister.Email == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request - Email wajib diisi saat pendaftaran",
			Data:    nil,
		})
	}
	if userRegister.Password == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request - Password wajib diisi saat pendaftaran",
			Data:    nil,
		})
	}

	var UserDB users.User
	UserDB.Name = userRegister.Name
	UserDB.Email = userRegister.Email
	UserDB.Password, _ = helper.Hash(userRegister.Password)
	UserDB.Address = userRegister.Address

	result := config.DB.Create(&UserDB)
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

func LoginUserController(c echo.Context) error {
	userLogin := users.UserLogin{}
	c.Bind(&userLogin)

	user := users.User{}
	result := config.DB.First(&user, "email = ?", userLogin.Email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusForbidden, response.BaseResponse{
				Code:    http.StatusForbidden,
				Message: "Email belum terdaftar",
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
	if !helper.CheckPasswordHash(userLogin.Password, user.Password) {
		return c.JSON(http.StatusForbidden, response.BaseResponse{
			Code:    http.StatusForbidden,
			Message: "Password tidak sesuai",
			Data:    nil,
		})
	}
	token, err := middlewares.GenerateTokenJWTUser(user.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Ada keselahan di server",
			Data:    nil,
		})
	}
	userResponse := users.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Hi,Welcome back",
		Data:    userResponse,
	})

}
