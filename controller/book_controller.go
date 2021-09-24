package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/config"
	books "project/model/Books"
	"project/model/loan"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func AddBookController(c echo.Context) error {
	//Build the request
	apiKey := "?key=" + "AIzaSyBkKjJlE2J3DvjifdHTWXr4JSLS6SRlcic"
	id := "pgjmDAAAQBAJ"
	request := "https://www.googleapis.com/books/v1/volumes/" + id + apiKey
	req, err := http.NewRequest("GET", request, nil)
	if err != nil {
		fmt.Println("Error is req: ", err)
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
	book.CopiesOwned = 3
	defer resp.Body.Close()

	result := config.DB.Create(&book)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create the data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Succes add book",
		"id":            book.Id,
		"title":         book.Title,
		"author":        book.Authors,
		"categories":    book.Categories,
		"publishedDate": book.PublishedDate,
		"imageLink":     book.Cover,
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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Buku ditemukan",
		"id":          book.Id,
		"title":       book.Title,
		"authors":     book.Authors,
		"categories":  book.Categories,
		"imageLinks":  book.Cover,
		"copiesOwned": book.CopiesOwned,
	})

}
func LoanBook(c echo.Context) error {
	//get data in database if exist
	var book books.Book
	err := config.DB.Where("id = ?", c.Param("id")).Find(&book).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "record not found",
		})
	}

	//Validate input
	var data loan.Loan
	data.UserID, _ = strconv.Atoi(c.Param("user_id"))
	data.BookID, _ = strconv.Atoi(c.Param("id"))

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
		book.CopiesOwned = book.CopiesOwned - 1
		fmt.Println(book.CopiesOwned)
		config.DB.Save(&book)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Succes to reserve book",
			"data":    data,
		})
	}

}
