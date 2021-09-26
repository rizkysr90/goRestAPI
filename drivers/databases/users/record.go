package users

import (
	"project/business/users"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        int    `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Name      string
	Address   string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *Users) ToDomain() users.User { //usecase akses db lewat domain,return domain interface untuk db adalah
	//struct domain object,data dari database harus dimaping kedalam struct
	//tersebut
	return users.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func FromDomain(domain users.User) Users { //domain akses db,jika domain ingin create record maka struct model di db harus di mapping ke struct domain
	return Users{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Address:   domain.Address,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
