package authentication

type LoggedInUser struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Firstname string `json:"first_name"`
	LastName  string `json:"last_name"`
	Token     string `json:"token"`
}
