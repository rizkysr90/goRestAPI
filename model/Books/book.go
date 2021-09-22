package book

type Book struct {
	Id      int    `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Authors string `json:"authors"`
	Cover   string `json:"cover"`
}
