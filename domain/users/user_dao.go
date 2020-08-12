package users

import (
	"github.com/trungkien71297/go_api_management/utils/errors"
	"net/http"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestError {
	if userDB[user.Id] != nil {
		return &errors.RestError{
			Message: "Existed",
			Code:    http.StatusBadGateway,
			Error:   "existed",
		}
	}
	userDB[user.Id] = user
	return nil
}

func Get(userId int64) (*User, *errors.RestError) {
	res := userDB[userId]

	if res == nil {
		return nil, &errors.RestError{}
	}
	return res, nil
}
