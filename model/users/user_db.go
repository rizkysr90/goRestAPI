package users

import (
	"project/model/loan"
	"time"
)

type User struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	CreatedAt time.Time
	Loan      []loan.Loan
}
