package admins

type AdminLogin struct {
	Email    string `form:"email"`
	Password int    `form:"password"`
}
