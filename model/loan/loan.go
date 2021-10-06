package loan

import (
	books "project/model/books"
	"project/model/status"
	"project/model/users"
	"time"
)

type Loan struct {
	Id         int           `gorm:"primaryKey" json:"id"`
	UserId     int           `json:"user_id"`
	User       users.User    `json:"user" gorm:"ForeignKey:UserId"`
	BookId     int           `json:"book_id"`
	Book       books.Book    `json:"book" gorm:"ForeignKey:BookId"`
	StatusId   int           `json:"status_id"`
	Status     status.Status `json:"status" gorm:"ForeignKey:StatusId"`
	LoanDate   time.Time     `json:"loan_date"`
	ReturnDate time.Time     `json:"return_date"`
	CreatedAt  time.Time     `json:"createdAt"`
	UpdatedAt  time.Time     `json:"updatedAt"`
}
