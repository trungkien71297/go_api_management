package services

import "github.com/trungkien71297/go_api_management/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
