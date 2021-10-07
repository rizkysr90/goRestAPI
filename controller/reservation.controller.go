package controller

import (
	"errors"
	"net/http"
	"project/config"
	books "project/model/Books"
	"project/model/loan"
	"project/model/response"
	"project/model/status"
	"project/model/users"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ReservationBookController(c echo.Context) error {
	var reservation loan.UserReservation
	var data loan.Loan
	var book books.Book
	var user users.User
	var code status.Status
	c.Bind(&reservation)
	err := config.DB.Where("id = ?", reservation.BookId).Find(&book).Error
	if err != nil {
		return c.JSON(http.StatusNoContent, response.BaseResponse{
			Code:    http.StatusNoContent,
			Message: "Buku tidak tersedia",
			Data:    nil,
		})
	}
	if book.Title == "" {
		return c.JSON(http.StatusNoContent, response.BaseResponse{
			Code:    http.StatusNoContent,
			Message: "Buku tidak tersedia",
			Data:    nil,
		})
	}

	data.BookId = reservation.BookId
	data.UserId = reservation.UserId
	data.StatusId = 1

	result := config.DB.Create(&data)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Status internal server error",
			Data:    nil,
		})
	}
	config.DB.Save(&data)
	config.DB.Preload("book").First(&book, "id = ?", reservation.BookId)
	config.DB.Preload("user").Find(&user, "id = ?", reservation.UserId)
	config.DB.Preload("status").Find(&code, "id = ?", data.StatusId)
	UserResponse := loan.UserReservationResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	resp := loan.ReservationResponse{
		Id:     data.Id,
		Status: code,
		User:   UserResponse,
		Book:   book,
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "OK - Pemesanan berhasil",
		Data:    resp,
	})

}
func ReservationProcces(c echo.Context) error {

	proccesReservation := loan.ProccesReservation{}
	c.Bind(&proccesReservation)
	loan := loan.Loan{}
	result := config.DB.First(&loan, "id = ?", proccesReservation.Id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, response.BaseResponse{
				Code:    http.StatusNoContent,
				Message: "Reservation Id tidak ditemukan",
				Data:    nil,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Internal Server Error",
				Data:    nil,
			})
		}
	}
	layout := "2006-01-02" //TEMPLATE PARSE STRING TO DATE
	loan.StatusId = proccesReservation.Status
	date, err := time.Parse(layout, proccesReservation.LoanDate)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    nil,
		})
	}

	loan.LoanDate = date
	res := config.DB.Save(&loan)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Pesanan dikonfirmasi",
		Data:    nil,
	})

}
func GetReservationById(c echo.Context) error {
	var data loan.Loan
	var user users.User
	var book books.Book
	var status status.Status
	c.Bind(&data)
	err := config.DB.Where("id = ?", data.Id).Find(&data).Error
	if err != nil {
		return c.JSON(http.StatusNoContent, response.BaseResponse{
			Code:    http.StatusNoContent,
			Message: "Reservation Id belum tersedia",
			Data:    nil,
		})
	}
	config.DB.Preload("book").First(&book, "id = ?", data.BookId)
	config.DB.Preload("user").First(&user, "id = ?", data.UserId)
	config.DB.Preload("status").First(&status, "id = ?", data.StatusId)
	UserResponse := loan.UserReservationResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	resp := loan.ReservationResponse{
		Id:     data.Id,
		Status: status,
		User:   UserResponse,
		Book:   book,
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Succes Retrieve Data",
		Data:    resp,
	})

}
