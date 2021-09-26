package users

//Usecase adalah tempat logic bussiness
import (
	"context"
	"errors"
)

type UserUseCase struct {
	Repo Repository
}

func NewUserUseCase(repo Repository) Usecase {
	return &UserUseCase{
		Repo: repo,
	}
}

func (uc *UserUseCase) LoginUser(ctx context.Context, email string, password string) (User, error) {
	if email == "" {
		return User{}, errors.New("email empty")
	}

	if password == "" {
		return User{}, errors.New("password empty")
	}

	user, err := uc.Repo.LoginUser(ctx, email, password)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

// func (uc *UserUseCase) RegisterUser(ctx context.Context, email string, password string, name string, address string) (User, error) {
// 	if email == "" {
// 		return User{}, errors.New("email empty")
// 	}
// 	if password == "" {
// 		return User{}, errors.New("password empty")
// 	}
// 	if name == "" {
// 		return User{}, errors.New("name empty")
// 	}
// 	user, err := uc.Repo.RegisterUser(ctx, email, password, name, address)

// 	if err != nil {
// 		return User{}, err
// 	}

// 	return user, nil
// }
