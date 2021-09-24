package books

import (
	"project/model/loan"
)

type Book struct {
	Id            int    `gorm:"primaryKey" json:"id"`
	Title         string `json:"title"`
	Authors       string `json:"authors"`
	Categories    string `json:"categories"`
	PublishedDate string `json:"publishedDate"`
	Cover         string `json:"cover"`
	CopiesOwned   int    `json:"copiesOwned"`
	Loan          []loan.Loan
}
