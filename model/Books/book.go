package books

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	Id            int            `gorm:"primaryKey" json:"id"`
	Title         string         `json:"title" gorm:"unique"`
	Authors       string         `json:"authors"`
	Categories    string         `json:"categories"`
	PublishedDate string         `json:"publishedDate"`
	Cover         string         `json:"cover"`
	CopiesOwned   int            `json:"copiesOwned"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
