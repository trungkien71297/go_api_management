package services

import (
	"github.com/trungkien71297/go_api_management/domain/users"
	"github.com/trungkien71297/go_api_management/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err

	}
	return &user, nil
}

func GetUsers(userId int64) (*users.User, *errors.RestError) {

	res, err := users.Get(userId)
	if err != nil {
		return nil, err
	}

	return res, nil

}
