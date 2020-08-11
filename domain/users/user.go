package users

type User struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"pwd"`
}
