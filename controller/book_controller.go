package controller

import (
	"encoding/json"
	"net/http"
	"project/config"
	"project/model/books"
	"project/model/loan"
	"project/model/response"
	"project/model/status"
	"project/model/users"
	"strings"

	"github.com/labstack/echo/v4"
)

func AddBookController(c echo.Context) error {
	//make request from google book api
	var call books.Calling
	c.Bind(&call)
	apiKey := "?key=" + "AIzaSyBkKjJlE2J3DvjifdHTWXr4JSLS6SRlcic"
	request := "https://www.googleapis.com/books/v1/volumes/" + call.VolumeUnique + apiKey
	req, err := http.NewRequest("GET", request, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Oops,Something wrong",
			Data:    nil,
		})
	}
	// create a Client
	client := &http.Client{}

	// Do sends an HTTP request and
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Oops,Something wrong",
			Data:    nil,
		})
	}
	var data books.GetBook

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Oops,Something wrong",
			Data:    nil,
		})
	}

	var book books.Book
	book.Authors = strings.Join(data.VolumeInfo.Authors, ",")
	book.Title = data.VolumeInfo.Title
	book.Cover = data.VolumeInfo.Cover.Medium
	book.Categories = strings.Join(data.VolumeInfo.Categories, ",")
	book.PublishedDate = data.VolumeInfo.PublishedDate
	book.CopiesOwned = call.Qty
	defer resp.Body.Close()

	result := config.DB.Create(&book)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Buku sudah tersedia sebelumnya",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Buku berhasil ditambahkan",
		Data:    book,
	})

}

func SearchBookByTitle(c echo.Context) error {
	q := c.QueryParam("title")
	target := "%" + q + "%"
	var book books.Book
	err := config.DB.Where("title LIKE ?", target).Find(&book).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error while search data",
		})
	}
	if book.Title == "" {
		return c.JSON(http.StatusAccepted, map[string]interface{}{
			"message": "Data buku tidak ditemukan",
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Data buku ditemukan",
		Data:    book,
	})

}
func LoanBook(c echo.Context) error {
	var reservation loan.UserReservation
	var data loan.Loan
	var book books.Book
	var user users.User
	var code status.Code
	c.Bind(&reservation)
	err := config.DB.Where("id = ?", reservation.BookId).Find(&book).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request - Buku tidak tersedia",
			Data:    nil,
		})
	}
	data.BookID = reservation.BookId
	// // err = config.DB.Where("id = ?", reservation.UserId).Find(&user).Error
	// // if err != nil {
	// // 	return c.JSON(http.StatusBadRequest, response.BaseResponse{
	// // 		Code:    http.StatusBadRequest,
	// // 		Message: "Bad Request - Akun belum terdaftar",
	// // 		Data:    nil,
	// // 	})
	// // }
	data.UserID = reservation.UserId
	data.CodeID = 1
	// err = config.DB.Where("id = ?", data.CodeID).Find(&status).Error
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, response.BaseResponse{
	// 		Code:    http.StatusInternalServerError,
	// 		Message: "Internal Server Error",
	// 		Data:    nil,
	// 	})
	// }

	result := config.DB.Create(&data)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to reserve book",
		})
	}
	config.DB.Save(&data)
	config.DB.Preload("Book").Find(&book)
	config.DB.Preload("User").Find(&user)
	config.DB.Preload("Code").Find(&code)
	// config.DB.Joins("JOIN books ON books.id = loans.book_id").Joins("JOIN users ON users.id = loans.user_id").Find(&data)
	UserResponse := loan.UserReservationResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	resp := loan.ReservationResponse{
		Id:   data.Id,
		Code: code,
		User: UserResponse,
		Book: book,
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "OK - Pemesanan berhasil",
		Data:    resp,
	})

}
