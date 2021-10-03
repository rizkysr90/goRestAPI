package status

type Status struct {
	Id      int    `gorm:"primaryKey" json:"id"`
	Message string `json:"message"`
}
