package services

import (
	"github.com/trungkien71297/go_api_management/domain/users"
	"github.com/trungkien71297/go_api_management/utils"
	"github.com/trungkien71297/go_api_management/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userServiceInterface interface {
	GetUser(userId int64) (*users.User, *errors.RestError)
	CreateUser(user users.User) (*users.User, *errors.RestError)
	GetAllUsers() ([]*users.User, *errors.RestError)
	DeleteUser(userId int64) (bool, *errors.RestError)
	EditUser(user *users.User) (*users.User, *errors.RestError)
}

type userService struct{}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestError) {

	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Password = utils.GetMD5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err

	}
	return &user, nil
}

func (s *userService) GetUser(userId int64) (*users.User, *errors.RestError) {

	res, err := users.Get(userId)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *userService) GetAllUsers() ([]*users.User, *errors.RestError) {
	res, err := users.GetAllUser()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *userService) DeleteUser(userId int64) (bool, *errors.RestError) {
	u := users.User{Id: userId}
	res, err := u.DeletetUser()
	if err != nil {
		//TODO error delete
	}
	return res, nil
}

func (s *userService) EditUser(user *users.User) (*users.User, *errors.RestError) {
	res, err := user.EditUser()
	if err != nil {
		//TODO error edit
	}
	if res {
		return user, nil
	}
	return nil, nil
}
