package controller

import (
	"encoding/json"
	"net/http"
	"project/config"
	"project/model/books"
	"project/model/response"
	"strings"

	"github.com/labstack/echo/v4"
)

func AddBookController(c echo.Context) error {
	//make request from google book api
	var call books.Calling
	c.Bind(&call)
	if call.VolumeUnique == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Required parameter",
			Data:    nil,
		})
	}
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
	if data.VolumeInfo.Title == "" {
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
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Status Internal Server Error",
			Data:    nil,
		})
	}
	if book.Title == "" {
		return c.JSON(http.StatusForbidden, response.BaseResponse{
			Code:    http.StatusForbidden,
			Message: "Data Buku Tidak Ditemukan",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    book,
	})

}
