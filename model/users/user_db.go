package users

import (
	"project/model/loan"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	Loan      []loan.Loan
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
