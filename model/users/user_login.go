package users

type UserLogin struct {
	Email    string `form:"email"`
	Password int    `form:"password"`
}
