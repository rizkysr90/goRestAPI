package book

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id            int    `gorm:"primaryKey" json:"id"`
	Title         string `json:"title"`
	Authors       string `json:"authors"`
	Categories    string `json:"categories"`
	PublishedDate string `json:"publishedDate"`
	Cover         string `json:"cover"`
	CopiesOwned   int
}
