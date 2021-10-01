package loan

import (
	books "project/model/books"
	"project/model/status"
	"project/model/users"
	"time"
)

type Loan struct {
	Id         int         `gorm:"primaryKey" json:"id"`
	UserID     int         `json:"user_id"`
	User       users.User  `json:"user"`
	BookID     int         `json:"book_id"`
	Book       books.Book  `json:"book"`
	CodeID     int         `json:"code_id"`
	Code       status.Code `json:"code"`
	LoanDate   time.Time   `json:"loan_date"`
	ReturnDate time.Time   `json:"return_date"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
}
