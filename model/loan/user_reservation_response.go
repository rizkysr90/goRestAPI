package loan

import (
	books "project/model/books"
	"project/model/status"
)

type ReservationResponse struct {
	Id     int                     `json:"id"`
	Status status.Status           `json:"status"`
	User   UserReservationResponse `json:"users"`
	Book   books.Book              `json:"book"`
}

type UserReservationResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
