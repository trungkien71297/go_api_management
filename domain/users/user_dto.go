package users

import (
	"net/http"
	"strings"

	"github.com/trungkien71297/go_api_management/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"pwd"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	var err errors.RestError

	if user.Email == "" {
		err = errors.RestError{
			Message: "Wrong type email",
			Code:    http.StatusBadRequest,
			Error:   "Email wrong",
		}
	}
	if err.Code == 0 {
		return nil
	}
	return &err
}
