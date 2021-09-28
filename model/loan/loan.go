package loan

import (
	"time"
)

type Loan struct {
	Id        int `gorm:"primaryKey" json:"id"`
	UserID    int
	BookID    int
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
