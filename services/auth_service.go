package services

import (
	"go_api_management/domain/authentication"
	"go_api_management/utils"
	"go_api_management/utils/errors"
)

var (
	AuthService authServiceInterface = &authService{}
)

type authServiceInterface interface {
	Login(username string, password string) (*authentication.LoggedInUser, *errors.RestError)
}

type authService struct{}

func (s *authService) Login(username string, password string) (*authentication.LoggedInUser, *errors.RestError) {
	var encodedPassword = utils.GetMD5(password)
	res, err := authentication.Login(username, encodedPassword)
	if err != nil {
		return nil, err
	}
	return res, nil
}
