package users

import (
	"context"
	"time"
)

//Membuat modelling domain bussiness ke dalam struct
type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Address  string
	Token    string
	// Loan      []loan.Loan
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface { //Menghubungkan entitas dengan usecase
	LoginUser(ctx context.Context, email string, password string) (User, error)
	// RegisterUser(ctx context.Context, email string, password string, name string, address string) (User, error)
}

type Repository interface { //Menghubungkan domain dengan DB,WEB dll
	LoginUser(ctx context.Context, email string, password string) (User, error)
	// RegisterUser(ctx context.Context, email string, password string, name string, address string) (User, error)
}
