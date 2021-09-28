package users

type UserLogin struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
