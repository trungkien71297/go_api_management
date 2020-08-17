package users

import (
	"encoding/json"
	"time"
)

type PublicUser struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	//	Firstname   string    `json:"first_name"`
	//	Lastname    string    `json:"last_name"`
	//	Email       string    `json:"email"`
	//	Password    string    `json:"pwd"`
	DateCreated time.Time `json:"date_created"`
}

type PrivateUser struct {
	Id          int64     `json:"id"`
	Username    string    `json:"username"`
	Firstname   string    `json:"first_name"`
	Lastname    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"pwd"`
	DateCreated time.Time `json:"date_created"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			Username:    user.Username,
			DateCreated: user.DateCreated,
		}
	}

	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
