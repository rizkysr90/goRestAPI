package loan

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	Id        int            `gorm:"primaryKey" json:"id"`
	UserID    int            `sql:"index"`
	BookID    int            `sql:"index"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
