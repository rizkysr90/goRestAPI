package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/config"
	book "project/model/Books"
	"strings"

	"github.com/labstack/echo/v4"
)

func AddBookController(c echo.Context) error {
	//Build the request
	apiKey := "?key=" + "AIzaSyBkKjJlE2J3DvjifdHTWXr4JSLS6SRlcic"
	id := "MOUREAAAQBAJ"
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
	var data book.GetBook

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}

	var book book.Book
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
