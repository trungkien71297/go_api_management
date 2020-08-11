package users

import "github.com/trungkien71297/go_api_management/utils/errors"

var (
	userDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestError {
	return nil
}

func Get(userId int64) (*User, *errors.RestError) {

	res := userDB[userId]
	_ = res

	return nil, nil
}
