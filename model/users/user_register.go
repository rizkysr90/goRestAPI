package users

type UserRegister struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Address  string `json:"address" form:"address"`
	Password string `json:"password" form:"password"`
}
