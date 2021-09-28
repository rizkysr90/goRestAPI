package controller

import (
	"errors"
	"fmt"
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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Nama tidak boleh kosong!",
		})
	}

	var UserDB users.User
	UserDB.Name = userRegister.Name
	UserDB.Email = userRegister.Email
	UserDB.Password, _ = helper.Hash(userRegister.Password)
	UserDB.Address = userRegister.Address

	result := config.DB.Create(&UserDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create the data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes create user",
		"id":      UserDB.Id,
		"name":    UserDB.Name,
	})
}

func LoginUserController(c echo.Context) error {
	userLogin := users.UserLogin{}
	c.Bind(&userLogin)

	user := users.User{}

	result := config.DB.First(&user, "email = ? AND password = ?", userLogin.Email, userLogin.Password)
	fmt.Println(userLogin.Email)
	fmt.Println(userLogin.Password)
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
		Message: "Berhasil login",
		Data:    userResponse,
	})

}
