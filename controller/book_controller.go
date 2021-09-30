package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/config"
	books "project/model/Books"
	"project/model/loan"
	"project/model/response"
	"project/model/users"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func AddBookController(c echo.Context) error {
	//make request from google book api
	bookID := c.FormValue("bookId")
	fmt.Println(bookID)
	apiKey := "?key=" + "AIzaSyBkKjJlE2J3DvjifdHTWXr4JSLS6SRlcic"
	request := "https://www.googleapis.com/books/v1/volumes/" + bookID + apiKey
	fmt.Println(request)
	req, err := http.NewRequest("GET", request, nil)
	if err != nil {
		fmt.Println("Error is : ", err)
	}
	// create a Client
	client := &http.Client{}

	// Do sends an HTTP request and
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in send req: ", err)
	}
	var data books.GetBook

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}

	var book books.Book
	book.Authors = strings.Join(data.VolumeInfo.Authors, ",")
	book.Title = data.VolumeInfo.Title
	book.Cover = data.VolumeInfo.Cover.Medium
	book.Categories = strings.Join(data.VolumeInfo.Categories, ",")
	book.PublishedDate = data.VolumeInfo.PublishedDate
	book.CopiesOwned, _ = strconv.Atoi(c.FormValue("qty"))
	defer resp.Body.Close()

	result := config.DB.Create(&book)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create the data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succes add book",
		"data":    book,
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
	//get data in database if exist
	var reservation loan.UserReservation
	c.Bind(&reservation)
	var book books.Book
	var user users.User
	err := config.DB.Where("id = ?", reservation.BookId).Find(&book).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusForbidden,
			Message: "Buku tidak tersedia",
			Data:    nil,
		})
	}
	err = config.DB.Where("id = ?", reservation.UserId).Find(&user).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusForbidden,
			Message: "Akun tidak terdaftar",
			Data:    nil,
		})
	}
	//Validate input
	var data loan.Loan
	data.UserID = reservation.UserId
	data.BookID = reservation.BookId
	data.Status = 0

	if book.CopiesOwned == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Book is not ready,stok = 0",
		})
	} else {
		result := config.DB.Create(&data)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to reserve book",
			})
		}

		config.DB.Save(&data)
		config.DB.Joins("JOIN users ON users.id = loans.user_id").Joins("JOIN books ON books.id = loans.book_id").Find(&data)
		UserResponse := loan.UserReservationResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}
		resp := loan.ReservationResponse{
			Id:          data.Id,
			StatusOrder: data.Status,
			User:        UserResponse,
			Book:        book,
		}

		return c.JSON(http.StatusOK, response.BaseResponse{
			Code:    http.StatusOK,
			Message: "Buku berhasil di pesan,tunggu konfirmasi pihak perpustakaan",
			Data:    resp,
		})
	}

}
