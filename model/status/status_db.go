package status

type Code struct {
	Id      int    `gorm:"primaryKey" json:"id"`
	Message string `json:"message"`
}
