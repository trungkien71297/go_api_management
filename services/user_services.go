package services

import (
	"github.com/trungkien71297/go_api_management/domain/users"
	"github.com/trungkien71297/go_api_management/utils"
	"github.com/trungkien71297/go_api_management/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {

	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Password = utils.GetMD5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err

	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestError) {

	res, err := users.Get(userId)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func GetAllUsers() ([]*users.User, *errors.RestError) {
	res, err := users.GetAllUser()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteUser(userId int64) (bool, *errors.RestError) {
	u := users.User{Id: userId}
	res, err := u.DeletetUser()
	if err != nil {
		//TODO error delete
	}
	return res, nil
}

func EditUser(user *users.User) (*users.User, *errors.RestError) {
	res, err := user.EditUser()
	if err != nil {
		//TODO error edit
	}
	if res {
		return user, nil
	}
	return nil, nil
}
