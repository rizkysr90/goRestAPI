package loan

import (
	books "project/model/Books"
	"project/model/users"
	"time"
)

type Loan struct {
	Id         int        `gorm:"primaryKey" json:"id"`
	UserID     int        `json:"user_id"`
	User       users.User `json:"user"`
	BookID     int        `json:"book_id"`
	Book       books.Book `json:"book"`
	Status     int        `json:"status"`
	LoanDate   time.Time  `json:"loan_date"`
	ReturnDate time.Time  `json:"return_date"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}
