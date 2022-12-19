package authentication

import (
	"go_api_management/datasources/mysql/user_db"
	"go_api_management/utils/errors"
)

const (
	loginQuery = "SELECT username, email, first_name, last_name FROM user_account where username=? and pwd=?"
)

func Login(username string, password string) (*LoggedInUser, *errors.RestError) {
	if err := user_db.Client.Ping(); err != nil {
		panic(err)
	}
	var user LoggedInUser
	st := user_db.Client.QueryRow(loginQuery, username, password)
	err := st.Scan(&user.Username, &user.Email, &user.Firstname, &user.LastName)
	if err != nil {
		return nil, &errors.RestError{
			Message: "Not found",
			Error:   "Not logged in",
			Code:    401,
		}
	}

	return &user, nil
}
