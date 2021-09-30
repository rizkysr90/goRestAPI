package admins

type AdminLogin struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
