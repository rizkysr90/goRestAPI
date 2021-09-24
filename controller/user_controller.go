package controller

import (
	"net/http"
	"project/config"
	"project/model/users"

	"github.com/labstack/echo/v4"
)

func RegisterController(c echo.Context) error {
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
	UserDB.Password = userRegister.Password
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

func LoginController(c echo.Context) error {
	var userLogin users.UserLogin
	var user users.User
	c.Bind(&userLogin)

	err := config.DB.Where("email = ? AND password = ?", userLogin.Email, userLogin.Password).Find(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get the data",
		})
	}
	if user.Name == "" {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": "user data tidak ditemukan,periksa kembali email / password anda",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"name":    user.Name,
	})

}
