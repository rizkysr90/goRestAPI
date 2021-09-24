package loan

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	Id       int `gorm:"primaryKey" json:"id"`
	UserID   int `sql:"index"`
	BookID   int `sql:"index"`
	LoanDate time.Time
}
