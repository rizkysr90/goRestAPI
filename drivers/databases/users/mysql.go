package users

import (
	"context"
	"project/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Login(ctx context.Context, email string, password string) (users.User, error) {
	var user Users
	result := rep.Conn.First(&user, "email = ? AND password = ?",
		email, password)

	if result.Error != nil {
		return users.User{}, result.Error
	}

	return user.ToDomain(), nil
}
