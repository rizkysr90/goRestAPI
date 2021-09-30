package loan

import (
	books "project/model/Books"
)

type ReservationResponse struct {
	Id          int                     `json:"id"`
	StatusOrder int                     `json:"status_order"`
	User        UserReservationResponse `json:"users"`
	Book        books.Book              `json:"book"`
}

type UserReservationResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
